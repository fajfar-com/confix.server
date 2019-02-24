package main

import (
	"github.com/fajfar-com/confix.server/internal/router"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	router.GetRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}