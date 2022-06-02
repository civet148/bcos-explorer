package bcos

import (
	"encoding/hex"
	"github.com/civet148/bcos-explorer/config"
	"github.com/civet148/log"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

//根据节点URL,链ID,分组ID和私钥等配置信息创建一个HTTP客户端对象用于部署合约、调用合约方法等操作
func NewBcosClient(cfg *config.Config, strPrivateKey string) (*client.Client, error) {

	privKey, err := hex.DecodeString(strPrivateKey)
	if err != nil {
		log.Errorf("decode private key [%s] error [%s]", strPrivateKey, err.Error())
		return nil, err
	}
	var bcosCfg = &conf.Config{
		IsHTTP:     true,
		NodeURL:    cfg.NodeUrl,
		PrivateKey: privKey,
		ChainID:    cfg.ChainID,
		GroupID:    cfg.GroupID,
	}
	return client.Dial(bcosCfg)
}