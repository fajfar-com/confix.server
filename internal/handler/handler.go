package handler

import "github.com/gin-gonic/gin"

func GetPing() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
