package models

type Eye struct {
	Id     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Errmsg string `gorm:"type:varchar(2000)" json:"errmsg"`
}

// 表示配置操作数据库的表名称
func (Eye) TableName() string {
	return "eye"
}
