package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// dsn := "fengfeng:root@tcp(127.0.0.1:5432)/db?charset=utf8mb4&parseTime=True&loc=Local"
	//DATABASE_URL=postgresql://postgres:rindexer@localhost:5440/postgres
	dsn := "host=127.0.0.1 user=postgres password=rindexer dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
