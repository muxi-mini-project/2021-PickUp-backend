package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrTokenInvalid(c *gin.Context, err error) {
	c.AbortWithStatusJSON(401, err)
}

func ErrBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(400, err)
}

func ErrServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(500, err)
}

func ErrUnauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(401, err)
}

type Error struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("Error(%s):%s.", err.ErrorCode, err.Message)
}
