package main

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/bcos-explorer/contract/token"
	"github.com/civet148/bcos-explorer/flags"
	"github.com/civet148/bcos-explorer/pkg/bcos"
	"github.com/civet148/bcos-explorer/pkg/utils"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	Version     = "v1.0.0"
	ProgramName = "bcos-invoker"
	FullName    = "bcos contract invoker"
)

//子命令
const (
	CmdNameDeploy   = "deploy"   //部署合约
	CmdNameBalance  = "balance"  //查询余额
	CmdNameTransfer = "transfer" //转账操作
)

const (
	ownerAddress       = "0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03" //合约owner账户地址
	ownerPrivateKey    = "3e5cd186c0de12c83fa4db6b6c5a93e64572721c4e262ce1498eaa2401658cf1" //合约owner账户私钥（所有需要上链的操作都需要私钥签名）
	receiverAddress =  "0x90Cfd4D61C9D4C63f2e4648229775ABa19ced8dF" //测试转账接收token的地址
)

func main() {
	subCmds := []*cli.Command{
		deployCmd, //子命令：deploy
		transferCmd, //子命令：transfer
	}
	app := &cli.App{
		Name:    ProgramName,
		Usage:   FullName,
		Version: Version,
		Flags: []cli.Flag{
		},
		Commands: subCmds,
		Action: nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}


//部署合约子命令
var deployCmd = &cli.Command{
	Name:      CmdNameDeploy,
	Usage:     "deploy contract",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     flags.CmdFlagNameNodeUrl,
			Usage:    "BCOS node url, eg. http://127.0.0.1:8545",
			Required: true, //必填：节点URL
		},
		&cli.Int64Flag{
			Name:  flags.CmdFlagNameChainID,
			Usage: "chain id",
			Value: 1, //非必填：chain id默认=1
		},
		&cli.IntFlag{
			Name:  flags.CmdFlagNameGroupID,
			Usage: "group id",
			Value: 1, //非必填：group id默认=1
		},
	},
	Action: func(cctx *cli.Context) error {

		cfg := &config.Config{
			NodeUrl: cctx.String(flags.CmdFlagNameNodeUrl), //从命令行参数获取节点URL
			ChainID: cctx.Int64(flags.CmdFlagNameChainID),  //从命令行参数获取分组ID（不传则默认=1）
			GroupID: cctx.Int(flags.CmdFlagNameGroupID),    //从命令行参数获取分组ID（不传则默认=1)
			PrivateKey: ownerPrivateKey, //私钥
		}

		client, err := bcos.NewBcosClient(cfg)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer client.Close()
		//部署代币合约并设置发行量为10000000(一千万)
		addr, tx, invoker, err := token.DeployToken(client.GetTransactOpts(), client, 10000000)
		if err != nil {
			log.Errorf("deploy error [%s]", err)
			return err
		}
		fmt.Printf("contract address: %s\n", addr.Hex()) //部署成功后的合约地址
		fmt.Printf("tx hash: %s\n", tx.Hash().Hex()) //部署成功的交易哈希(可通过浏览器查询)

		//查询合约拥有者地址余额（合约刚刚部署成功的状态下余额应为一千万)
		var balance uint64
		owner, _ := utils.EtherAddr(ownerAddress) //将字符串类型的地址转成以太坊地址结构

		//查询类的合约方法不需要签名，因此需要的是CallOpts对象而不是TransactOpts
		balance, err = invoker.BalanceOf(client.GetCallOpts(), owner)
		if err != nil {
			log.Errorf("call BalanceOf error [%s]", err)
			return err
		}
		//部署合约的账户就是官方账户，当前余额一千万(跟token发行量保持一致)
		fmt.Printf("owner %s balance %d tokens\n", owner.String(), balance)
		return nil
	},
}


//转账子命令
var transferCmd = &cli.Command{
	Name:      CmdNameTransfer,
	Usage:     "transfer token to other address",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     flags.CmdFlagNameNodeUrl,
			Usage:    "BCOS node url, eg. http://127.0.0.1:8545",
			Required: true, //必填：节点URL
		},
		&cli.StringFlag{
			Name:     flags.CmdFlagNameContractAddr,
			Usage:    "BCOS contract address",
			Required: true, //必填：合约地址
		},
		&cli.Int64Flag{
			Name:  flags.CmdFlagNameChainID,
			Usage: "chain id",
			Value: 1, //非必填：chain id默认=1
		},
		&cli.IntFlag{
			Name:  flags.CmdFlagNameGroupID,
			Usage: "group id",
			Value: 1, //非必填：group id默认=1
		},
	},
	Action: func(cctx *cli.Context) error {

		cfg := &config.Config{
			NodeUrl: cctx.String(flags.CmdFlagNameNodeUrl), //从命令行参数获取节点URL
			ChainID: cctx.Int64(flags.CmdFlagNameChainID),  //从命令行参数获取分组ID（不传则默认=1）
			GroupID: cctx.Int(flags.CmdFlagNameGroupID),    //从命令行参数获取分组ID（不传则默认=1)
			PrivateKey: ownerPrivateKey, //私钥
		}

		client, err := bcos.NewBcosClient(cfg)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer client.Close()
		strContractAddr := cctx.String(flags.CmdFlagNameContractAddr)
		ca, _ := utils.EtherAddr(strContractAddr) //合约地址

		var invoker *token.Token
		invoker , err = token.NewToken(ca, client) //创建一个合约调用对象
		if err != nil {
			log.Errorf("new token invoker error [%s]", err)
			return err
		}
		var balance uint64
		owner, _ := utils.EtherAddr(ownerAddress) //将字符串类型的地址转成以太坊地址结构
		receiver, _ := utils.EtherAddr(receiverAddress) //将字符串类型的地址转成以太坊地址结构
		tx, receipt, err := invoker.TransferFrom(client.GetTransactOpts(), owner, receiver, 100) //从owner地址给接收者地址转账100个token
		if err != nil {
			log.Errorf("transfer token error [%s]", err)
			return err
		}

		fmt.Printf("tx hash [%s]\n", tx.Hash().String()) //转账调用成功的交易哈希(有交易哈希并不代表合约成功执行，仅上链成功，需判断回执状态)

		//如果交易回执Status!=0表示执行合约方法失败
		if receipt.Status != types.Success {
			return log.Errorf("transfer tx failed, code [%v] message [%s]", receipt.Status, receipt.GetErrorMessage())
		}

		//转账后查询合约拥有者地址余额
		balance, err = invoker.BalanceOf(client.GetCallOpts(), owner)
		if err != nil {
			log.Errorf("call BalanceOf error [%s]", err)
			return err
		}
		fmt.Printf("owner %s balance %d tokens\n", owner.String(), balance)
		//转账后查询接收人地址余额
		balance, err = invoker.BalanceOf(client.GetCallOpts(), receiver)
		if err != nil {
			log.Errorf("call BalanceOf error [%s]", err)
			return err
		}
		//部署合约的账户就是官方账户，当前余额一千万(跟token发行量保持一致)
		fmt.Printf("receiver %s balance %d tokens\n", receiver.String(), balance)
		return nil
	},
}
