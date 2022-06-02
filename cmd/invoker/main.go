package main

import (
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/bcos-explorer/pkg/bcos"
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
	CmdFlagNameNodeUrl      = "node-url"      //节点URL参数名
	CmdFlagNameChainID      = "chain-id"      //链ID(默认1)
	CmdFlagNameGroupID      = "group-id"      //分组ID(默认1)
	CmdFlagNameContractAddr = "contract-addr" //合约地址
	CmdFlagNameAccountAddr  = "account-addr"  //账户地址
)

const (
	ownerAddress       = "0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03" //合约owner账户地址
	ownerPrivateKey    = "3e5cd186c0de12c83fa4db6b6c5a93e64572721c4e262ce1498eaa2401658cf1" //合约owner账户私钥（所有需要上链的操作都需要私钥签名）
)

func main() {
	subCmds := []*cli.Command{
		deployCmd, //子命令：deploy
	}
	app := &cli.App{
		Name:    ProgramName,
		Usage:   FullName,
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     CmdFlagNameNodeUrl,
				Required: true, //节点地址参数必填
			},
			&cli.StringFlag{
				Name:     CmdFlagNameContractAddr,
				Required: true, //合约地址参数必填
			},
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
			Name:     CmdFlagNameNodeUrl,
			Usage:    "BCOS node url",
			Required: true, //必填：节点URL
		},
		&cli.Int64Flag{
			Name:  CmdFlagNameChainID,
			Usage: "chain id",
			Value: 1, //非必填：chain id默认=1
		},
		&cli.IntFlag{
			Name:  CmdFlagNameGroupID,
			Usage: "group id",
			Value: 1, //非必填：group id默认=1
		},
	},
	Action: func(cctx *cli.Context) error {

		cfg := &config.Config{
			NodeUrl: cctx.String(CmdFlagNameNodeUrl), //从命令行参数获取节点URL
			ChainID: cctx.Int64(CmdFlagNameChainID),  //从命令行参数获取分组ID（不传则默认=1）
			GroupID: cctx.Int(CmdFlagNameGroupID),    //从命令行参数获取分组ID（不传则默认=1)
		}

		client, err := bcos.NewBcosClient(cfg, ownerPrivateKey)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer client.Close()

		//addr, tx, _, err := printer.DeployNftPrinter(client.GetTransactOpts(), client)
		//if err != nil {
		//	log.Errorf("deploy error [%s]", err)
		//	return err
		//}
		//log.Infof("contract address: %s", addr.Hex()) // the address should be saved
		//log.Infof("contract deploy tx hash: %s", tx.Hash().Hex())
		return nil
	},
}
