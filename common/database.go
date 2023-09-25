package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "20222022Ynn."
	charset := "utf8"
	//root:20222022Ynn.@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fialed to connect database err: " + err.Error())
	}
	db.AutoMigrate(&model.User{}) //自动创建数据表

	return db
}

// GetDB 定义一个方法来获取
func GetDB() *gorm.DB {
	return DB
}
