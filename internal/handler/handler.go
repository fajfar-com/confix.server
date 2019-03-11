package handler

import (
	"github.com/fajfar-com/confix.server/internal/configMng"
	"github.com/gin-gonic/gin"
)

var configManager configMng.ConfigurationManager

func init(){
	configManager = configMng.GetConfigurationManager()
}

func GetPing() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func GetApplicationConfiguration() gin.HandlerFunc {
	return func(context *gin.Context) {
		var appName = context.Param("app")
		var app = configManager.GetApplication(appName)

		context.JSON(200, app.Configs)
	}
}