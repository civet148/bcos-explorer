package service

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/bcos-explorer/flags"
	"github.com/civet148/bcos-explorer/pkg/bcos"
	"github.com/civet148/bcos-explorer/types"
	"github.com/civet148/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	btypes "github.com/FISCO-BCOS/go-sdk/core/types"
)

const(
	ZERO_ADDR = "0x0000000000000000000000000000000000000000"
)

type Explorer struct {
	cfg *config.Config
}

func NewExplorer(cfg *config.Config) *Explorer {

	return &Explorer{
		cfg: cfg,
	}
}

func (m *Explorer) Run(cctx *cli.Context) error {
	var ctx = context.Background()
	//创建HTTP客户端
	client, err := bcos.NewBcosClient(m.cfg)
	if err != nil {
		return log.Errorf("new bcos client error [%s]", err)
	}

	var height int64
	if cctx.IsSet(flags.CmdFlagNameHeight) {//如果通过命令行参数指定高度则使用传入的值
		height = cctx.Int64(flags.CmdFlagNameHeight)
	} else {
		height, err = client.GetBlockNumber(ctx) //查询当前BCOS区块链最新高度
		if err != nil {
			return log.Errorf("get block height error [%s]", err)
		}
	}

	log.Infof("block height [%d] handling...", height)

	//解析最新高度的区块数据
	var block []byte
	block, err = client.GetBlockByNumber(ctx, height, true)
	if err != nil {
		return log.Errorf("get block by height error [%s]", err)
	}
	//log.Infof("height [%d] block data [%s]", height, block)
	header := &types.BlockHeader{}
	//返回的block数据就是一段JSON文本的二进制数据
	if err = json.Unmarshal(block, header); err != nil {
		return log.Errorf("unmarshal tx data error [%s]", err)
	}
	log.Json(header)

	//通过contract/token/Token.abi文件解析合约调用的方法名、参数以及值（类似以太坊浏览器etherscan.io看到的结果)
	for _, tx := range header.Transactions {
		//当to地址是0x0000000000000000000000000000000000000000时表示当次交易为部署合约(一般不需要过多解析)
		if tx.To != ZERO_ADDR {
			var hash common.Hash
			hash = common.HexToHash(tx.Hash)

			//解析合约方法transferFrom调用参数名和值
			if err = m.parseTxMethod(header, tx); err != nil {
				return log.Errorf(err.Error())
			}

			//查询交易回执
			receipt, err := client.GetTransactionReceipt(ctx, hash)
			if err != nil {
				return log.Errorf("get transaction receipt error [%s]", err)
			}
			fmt.Printf("\n----------------------------------------------------\n")
			if receipt.Status != btypes.Success {//交易不成功
				fmt.Printf("tx receipt status error message [%s]\n", receipt.GetErrorMessage())
			} else {//交易成功
				fmt.Printf("tx receipt status OK\n")
			}
		} else {
			fmt.Printf("height [%v] tx hash [%s] is a deploy transaction\n", height, tx.Hash) //合约部署交易
		}
	}
	return nil
}

func (m *Explorer) parseTxMethod(header *types.BlockHeader, tx types.TransactionInfo) error {
	//ABI信息从Token.abi文件中复制到这里，后续可以自行改成通过指定文件形式加载
	var strTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint64\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"amount\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	var abiMsg abi.ABI

	err := abiMsg.UnmarshalJSON([]byte(strTokenABI))
	if err != nil {
		return log.Errorf("Unmarshal abi content error [%s]", err.Error())
	}

	//截取input数据第2-10位置的方法ID
	methodSigData, err := hex.DecodeString(tx.Input[2:10])
	if err != nil {
		return log.Errorf(err.Error())
	}
	method, err := abiMsg.MethodByID(methodSigData)
	if err != nil {
		return log.Errorf(err.Error())
	}
	fmt.Printf("method name [%s]\n", method.Name) //打印合约方法名
	for _, v := range method.Inputs {
		fmt.Printf("input type [%s] name [%s] \n", v.Type, v.Name) //打印参数类型和名称
	}
	// 获取合约方法的值(从第10字节开始)
	values := make(map[string]interface{})
	valueBytes, err := hex.DecodeString(tx.Input[10:])
	if err != nil {
		return log.Errorf(err.Error())
	}
	err = method.Inputs.UnpackIntoMap(values, valueBytes)
	if err != nil {
		return log.Errorf(err.Error())
	}
	for k, v := range values {
		fmt.Printf("input name [%s] value [%v] \n", k, v) //打印参数名称和值
	}
	return nil
}