package main

import (
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	// oracle://user:password@127.0.0.1:1521/service
	url := oracle.BuildUrl("127.0.0.1", 1521, "orcl", "snow", "oracle", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		// panic error or log error info
		panic("failed to connect to Oracle")
	}
	// do somethings

	//迁移schema
	// db.AutoMigrate(&Product{})

	// //Create
	// db.Create(&Product{Code: "D42", Price: 100})
	// // Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//db.Delete(&product, 1)


}