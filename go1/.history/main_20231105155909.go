package main

import (
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

func main() {
	// oracle://user:password@127.0.0.1:1521/service
	url := oracle.BuildUrl("127.0.0.1", 1521, "service", "snow", "oracle", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		// panic error or log error info
		panic("failed to connect to Oracle")
	}
	defer db.Close(
	// do somethings
}