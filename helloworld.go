package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", sayHelloToTheWorld)

	router.Run()
}

func sayHelloToTheWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
