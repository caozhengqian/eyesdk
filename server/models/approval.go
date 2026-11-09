package models

type Approval struct {
	RindexerId      int    `json:"rindexerId"`
	ContractAddress string `json:"contractAddress"`
	Owner           string `gorm:"column:_owner" json:"owner"`
	Spender         string `gorm:"column:_spender" json:"spender"`
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
func (Approval) TableName() string {
	return "approval"
}
