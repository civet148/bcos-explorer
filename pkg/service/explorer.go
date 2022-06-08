package service

import (
	"context"
	"encoding/json"
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/bcos-explorer/pkg/bcos"
	"github.com/civet148/bcos-explorer/types"
	"github.com/civet148/log"
)

type Explorer struct {
	cfg *config.Config
}

func NewExplorer(cfg *config.Config) *Explorer {

	return &Explorer{
		cfg: cfg,
	}
}

func (m *Explorer) Run() error {
	var ctx = context.Background()
	//创建HTTP客户端
	client, err := bcos.NewBcosClient(m.cfg)
	if err != nil {
		return log.Errorf("new bcos client error [%s]", err)
	}

	var height int64
	height, err = client.GetBlockNumber(ctx) //查询当前BCOS区块链最新高度
	if err != nil {
		return log.Errorf("get block height error [%s]", err)
	}
	log.Infof("latest block height [%d]", height)

	//解析最新高度的区块数据
	var block []byte
	block, err = client.GetBlockByNumber(ctx, height, true)
	if err != nil {
		return log.Errorf("get block by height error [%s]", err)
	}
	//log.Infof("height [%d] block data [%s]", height, block)
	tx := &types.BlockHeader{}
	//返回的block数据就是一段JSON文本的二进制数据
	if err = json.Unmarshal(block, tx); err != nil {
		return log.Errorf("unmarshal tx data error [%s]", err)
	}
	log.Json(tx)


	return nil
}
