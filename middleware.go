package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	fmt.Println("Received " + c.Request.Method + " request at " + c.Request.RequestURI)
	c.Next()
}

