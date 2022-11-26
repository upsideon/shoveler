package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := initializeRouter()
	router.Run(":8080")
}

func initializeRouter() *gin.Engine {
	router := gin.Default()
	return router
}
