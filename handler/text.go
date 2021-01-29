package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Text(context *gin.Context) {
	fmt.Println("qqqq:", context.FullPath())
	context.Writer.Write([]byte("hello\n"))
}
