package main

import (
	"log"
	"pickup/model"
	"pickup/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @title pick up
// @version 1.0
// @description pickup 顺风车

// @host 39.102.42.156
// @BasePath /pickup

// @Schemas http
func main() {
	r := gin.Default()

	model.Db.Init()
	defer model.Db.Close()

	router.Router(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
