package main

import (
	"fmt"
	"github.com/civet148/go-wallet"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	Version     = "v1.0.0"
	ProgramName = "bcos-account"
	FullName    = "BCOS account generator"
)

const (
	CmdNameNew     = "new"
	CmdNameRecover = "recover"
)

func main() {
	subCmds := []*cli.Command{
		newCmd,     //子命令：新创建账户
		recoverCmd, //子命令：通过助记词恢复公/私钥信息
	}
	app := &cli.App{
		Name:     ProgramName,
		Usage:    FullName,
		Version:  Version,
		Flags:    []cli.Flag{},
		Commands: subCmds,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

//创建账户密钥
var newCmd = &cli.Command{
	Name:      CmdNameNew,
	Usage:     "new bcos/ethereum account",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		var wc = wallet.NewWalletEthereum(wallet.OpType_Create)
		fmt.Printf("address [%s]\n", wc.GetAddress())        //公钥地址
		fmt.Printf("public key [%s]\n", wc.GetPublicKey())   //公钥
		fmt.Printf("private key [%s]\n", wc.GetPrivateKey()) //私钥
		fmt.Printf("phrase [%s]\n", wc.GetPhrase())          //助记词(用于恢复私钥)
		return nil
	},
}

//通过助记词恢复用户密钥
var recoverCmd = &cli.Command{
	Name:      CmdNameRecover,
	Usage:     "recover bcos/ethereum account from phrase",
	ArgsUsage: "",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		strPhrase := cctx.Args().First()
		if strPhrase == "" {
			fmt.Printf("请输入助记词\n")
			return nil
		}
		var wc = wallet.NewWalletEthereum(wallet.OpType_Recover, strPhrase)
		fmt.Printf("address [%s]\n", wc.GetAddress())        //公钥地址
		fmt.Printf("public key [%s]\n", wc.GetPublicKey())   //公钥
		fmt.Printf("private key [%s]\n", wc.GetPrivateKey()) //私钥
		fmt.Printf("phrase [%s]\n", wc.GetPhrase())          //助记词(用于恢复私钥)
		return nil
	},
}
