package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect()  {
	d,err:=gorm.Open("mysql","surbhi:Surbhi1611@(localhost:3306)/BookStore?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	db=d
}

func GetDB()  *gorm.DB{
	return db
}