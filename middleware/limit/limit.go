package middleware

import (
	"github.com/MorningSong/go-gin/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
	"net/http"
	"time"
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
