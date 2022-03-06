package config

import (
	"gorm.io/gorm"
)



var (
	db *gorm.DB
)


func Connect(){
	d, err := gorm.Open("mysql","data2020:Data2020!/bookstorecharset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB {
	return db
}