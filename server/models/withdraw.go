package models

type Withdraw struct {
	RindexerId      int    `json:"rindexerId"`
	ContractAddress string `json:"contractAddress"`
	Token           string `json:"token"`
	User            string `json:"user"`
	Amount          string `json:"amount"`
	Balance         string `json:"balance"`
	TxHash          string `json:"txHash"`
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  string `json:"blockTimestamp"`
	BlockHash       string `json:"blockHash"`
	Network         string `json:"network"`
	TxIndex         string `json:"txIndex"`
	LogIndex        string `json:"logIndex"`
}

// 表示配置操作数据库的表名称
func (Withdraw) TableName() string {
	return "withdraw"
}
