package Dao

import (
	"blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB*gorm.DB
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	username := "root"
	password := "512612lj"
	database := "gin"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db ,err:= gorm.Open(driverName, args)
	if err!=nil {
		fmt.Println("failed connect database",err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}

func GetDB()*gorm.DB  {
	return DB
}