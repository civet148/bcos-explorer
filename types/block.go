package types


type BlockHeader struct {
	DbHash           string            `json:"dbHash"`
	ExtraData        []interface{}     `json:"extraData"`
	GasLimit         string            `json:"gasLimit"`
	GasUsed          string            `json:"gasUsed"`
	Hash             string            `json:"hash"`
	LogsBloom        string            `json:"logsBloom"`
	Number           string            `json:"number"`
	ParentHash       string            `json:"parentHash"`
	ReceiptsRoot     string            `json:"receiptsRoot"`
	Sealer           string            `json:"sealer"`
	SealerList       []string          `json:"sealerList"`
	SignatureList    []SignatureInfo   `json:"signatureList"`
	StateRoot        string            `json:"stateRoot"`
	Timestamp        string            `json:"timestamp"`
	Transactions     []TransactionInfo `json:"transactions"`
	TransactionsRoot string            `json:"transactionsRoot"`
}

type SignatureInfo struct {
	Index     string  `json:"index"`
	Signature string `json:"signature"`
}

type TransactionInfo struct {
	BlockHash   string `json:"blockHash"`
	BlockLimit  string `json:"blockLimit"`
	BlockNumber string `json:"blockNumber"`
	ChainId     string `json:"chainId"`
	ExtraData   string `json:"extraData"`
	From        string `json:"from"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	GroupId     string `json:"groupId"`
	Hash        string `json:"hash"`
	Input       string `json:"input"`
	Nonce       string `json:"nonce"`
	Signature   struct {
		R         string `json:"r"`
		S         string `json:"s"`
		Signature string `json:"signature"`
		V         string `json:"v"`
	} `json:"signature"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
}
