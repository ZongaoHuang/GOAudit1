package main

import (
    "github.com/jinzhu/gorm"
    "github.com/mattn/go-oci8"

)

func mn() {

db, err := gorm.Open(“oci8”, “/@:/”)

if err != nil {

panic(err)

}

defer db.Close()

}

func main(){


}