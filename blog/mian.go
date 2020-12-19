package main

import (
	"blog/Dao"
	"blog/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	db := Dao.InitDB()
	defer db.Close()
	engine := gin.Default()
	engine = router.CollectRouter(engine)
	engine.Run(":8081")
}

