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
	CmdFlagNameNodeUrl = "node-url" //节点URL参数名
)

const (
	DefaultNodeUrl = "http://192.168.20.108:8545" //默认节点
)

func main() {
	app := &cli.App{
		Name:    ProgramName,
		Usage:   FullName,
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  CmdFlagNameNodeUrl,
				Value: DefaultNodeUrl,
			},
		},
		Commands: nil,
		Action: func(cctx *cli.Context) error {
			cfg := &config.Config{
				NodeUrl: cctx.String(CmdFlagNameNodeUrl),
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
