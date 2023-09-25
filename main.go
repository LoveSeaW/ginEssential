package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = CollectRouters(r)

	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
