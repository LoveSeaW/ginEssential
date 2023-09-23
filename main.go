package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(20);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

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
	db.AutoMigrate(&User{}) //自动创建数据表

	return db
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		//如果名称没有传，就给一个10位的字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, telephone, password)

		//判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户存在不允许注册"})
			return
		}

		//创建用户
		newUser := User{
			Model:     gorm.Model{},
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		//返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDLFJASDOLFJQOPAF")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
