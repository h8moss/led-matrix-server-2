package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
  fmt.Println("Initializing server...")
  router := gin.Default()
  router.Use(static.Serve("/", static.LocalFile("./frontend/dist/index.html", false)))

  // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/test", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }

  fmt.Println("Listening on port 8080")
  router.Run(":8080")
}

