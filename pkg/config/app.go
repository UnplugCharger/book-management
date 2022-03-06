package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)



var (
	db *gorm.DB
)


func Connect(){
	//Username and password for mysql
	d, err := gorm.Open("mysql","data2020:Data2020!@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
    
	if err != nil {
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB {
	return db
}