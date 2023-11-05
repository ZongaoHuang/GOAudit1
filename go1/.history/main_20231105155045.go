package main

import (
	oracle "github.com/godoes/gorm-oracle"
	go_ora "github.com/sijms/go-ora/v2"
	"gorm.io/gorm"
)

func main() {
	// oracle://user:password@127.0.0.1:1521/service
	url := go_ora.BuildUrl("127.0.0.1", int("1521"), "service", "snow", "oracle", map[string]string(nil))
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		// panic error or log error info
	}

	// do somethings
}
