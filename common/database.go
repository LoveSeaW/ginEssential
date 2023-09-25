package common

import (
	"fmt"
	"ginEssential/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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

	DB = db

	return db
}

// GetDB 定义一个方法来获取
func GetDB() *gorm.DB {
	return DB
}
