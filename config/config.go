package config

type Config struct {
	NodeUrl string `json:"node_url" toml:"node_url"` //节点URL
	Height  string `json:"height" toml:"height"`     //区块高度
	ChainID int64  `json:"chain_id" toml:"chain_id"` //链ID
	GroupID int    `json:"group_id" toml:"group_id"` //分组ID
}
