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
	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{IDCode: "D42", Price: 100})


}