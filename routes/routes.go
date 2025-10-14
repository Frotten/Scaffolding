package routes

import (
	"WebApp/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/Done", func(c *gin.Context) {
		fmt.Println("Done")
		c.String(http.StatusOK, "Well done")
	})
	zap.L().Info("SetUp Server ...")
	return r
}
