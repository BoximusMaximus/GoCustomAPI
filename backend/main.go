package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	// create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  println("Listening on port 8080")
  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run()

}	