package models

type Transfer struct {
	RindexerId      int    `json:"rindexerId"`
	ContractAddress string `json:"contractAddress"`
	From            string `gorm:"column:_from" json:"from"`
	To              string `gorm:"column:_to" json:"to"`
	Value           string `gorm:"column:_value" json:"value"`
	TxHash          string `json:"txHash"`
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  string `json:"blockTimestamp"`
	BlockHash       string `json:"blockHash"`
	Network         string `json:"network"`
	TxIndex         string `json:"txIndex"`
	LogIndex        string `json:"logIndex"`
}

// 表示配置操作数据库的表名称
func (Transfer) TableName() string {
	return "transfer"
}
