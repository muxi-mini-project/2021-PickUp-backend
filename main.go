package main

import (
	"log"
	"pickup/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	router.Router(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
