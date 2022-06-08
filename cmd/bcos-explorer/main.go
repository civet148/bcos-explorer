package main

import (
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/bcos-explorer/pkg/service"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	Version     = "v1.0.0"
	ProgramName = "bcos-explorer"
	FullName    = "BCOS block explorer"
)

const (
	CmdFlagNameNodeUrl      = "node-url"      //节点URL参数名
	CmdFlagNameChainID      = "chain-id"      //链ID(默认1)
	CmdFlagNameGroupID      = "group-id"      //分组ID(默认1)
	CmdFlagNameContractAddr = "contract-addr" //合约地址
	CmdFlagNamePrivateKey   = "private-key"   //私钥
)

const (
	defaultAddress    = "0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03"                       //合约owner账户地址
	defaultPrivateKey = "3e5cd186c0de12c83fa4db6b6c5a93e64572721c4e262ce1498eaa2401658cf1" //合约owner账户私钥（所有需要上链的操作都需要私钥签名）
)

func main() {
	app := &cli.App{
		Name:    ProgramName,
		Usage:   FullName,
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     CmdFlagNameNodeUrl,
				Required: true,
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
			&cli.StringFlag{
				Name:  CmdFlagNamePrivateKey,
				Usage: "BCOS account private key",
				Value: defaultPrivateKey, //默认值
			},
			&cli.StringFlag{
				Name:     CmdFlagNameContractAddr,
				Usage:    "BCOS contract address",
				Required: true, //必填：合约地址
			},
		},
		Commands: nil,
		Action: func(cctx *cli.Context) error {

			cfg := &config.Config{
				NodeUrl:      cctx.String(CmdFlagNameNodeUrl),      //BCOS节点URL
				ChainID:      cctx.Int64(CmdFlagNameChainID),       //从命令行参数获取分组ID（不传则默认=1）
				GroupID:      cctx.Int(CmdFlagNameGroupID),         //从命令行参数获取分组ID（不传则默认=1)
				ContractAddr: cctx.String(CmdFlagNameContractAddr), //合约地址
				PrivateKey:   cctx.String(CmdFlagNamePrivateKey),   //私钥
			}
			s := service.NewExplorer(cfg)
			return s.Run()
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}
