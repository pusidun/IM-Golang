package dao

import (
	"fmt"
	"im-golang/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "192.168.19.136"
	port := "3306"
	database := "im_db"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		fmt.Println(err)
		panic("failed to open db")
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func DBConn() *gorm.DB {
	return DB
}
