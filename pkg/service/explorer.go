package service

import "github.com/civet148/bcos-explorer/config"

type Explorer struct {
	cfg *config.Config
}

func NewExplorer(cfg *config.Config) *Explorer {

	return &Explorer{
		cfg: cfg,
	}
}

func (m *Explorer) Run() error {

	return nil
}