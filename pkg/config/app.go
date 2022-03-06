package config

import (
	"gorm.io/gorm"
)



var (
	db *gorm.DB
)


func Connect(){
	d, err := gorm.Open("mysql","data2020:Data2020!/bookstore")

	if err != nil {
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB {
	return db
}