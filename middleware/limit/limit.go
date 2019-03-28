package middleware

import (
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/MorningSong/go-gin/pkg/e"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

var (
	rl = simpleratelimit.New(100000, time.Minute)
)

// 中间件，用令牌桶限制请求频率
func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		if rl.Limit() {
			code = e.FORBIDDEN
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusForbidden, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
