package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UsersComment(c *gin.Context) {
	uid := c.Param("uid")
	comment_id := c.Query("comment_id")

	fmt.Println(uid, comment_id)
}
