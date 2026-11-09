package models

type Order struct {
	RindexerId      int    `json:"rindexerId"`
	ContractAddress string `json:"contractAddress"`
	Id              string `json:"id"`
	User            string `json:"user"`
	TokenGet        string `json:"tokenGet"`
	AmountGet       string `json:"amountGet"`
	TokenGive       string `json:"tokenGive"`
	AmountGive      string `json:"amountGive"`
	Timestamp       string `json:"timestamp"`
	TxHash          string `json:"txHash"`
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  string `json:"blockTimestamp"`
	BlockHash       string `json:"blockHash"`
	Network         string `json:"network"`
	TxIndex         string `json:"txIndex"`
	LogIndex        string `json:"logIndex"`
}

// 表示配置操作数据库的表名称
func (Order) TableName() string {
	return "order"
}
