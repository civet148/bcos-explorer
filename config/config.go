package config

type Config struct {
	NodeUrl      string `json:"node_url" toml:"node_url" cli:"node_url"`                //节点URL
	Height       int64 `json:"height" toml:"height" cli:"height"`                      //区块高度
	ChainID      int64  `json:"chain_id" toml:"chain_id" cli:"chain_id"`                //链ID
	GroupID      int    `json:"group_id" toml:"group_id" cli:"group_id"`                //分组ID
	Address      string `json:"address" toml:"address" cli:"address"`                   //公钥地址
	PrivateKey   string `json:"private_key" toml:"private_key" cli:"private_key"`       //私钥
	ContractAddr string `json:"contract_addr" toml:"contract_addr" cli:"contract_addr"` //合约地址
}
