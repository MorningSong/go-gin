package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/MorningSong/go-gin/pkg/app"
	"github.com/MorningSong/go-gin/pkg/e"
)

func HealthCheck(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"status": "UP",
	})
}

