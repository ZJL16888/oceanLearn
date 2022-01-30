package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oceanLearn/util"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Fail(c, nil, fmt.Sprint(err))
			}
		}()

		c.Next()
	}
}
