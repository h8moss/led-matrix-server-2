package main

import (
	"fmt"
  "example.com/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
  fmt.Println("Initializing server...")
  router := gin.Default()
  router.Use(static.Serve("/", static.LocalFile("./frontend/dist/index.html", false)))

  // Setup route group for the API
  api.CreateApi(router)
  fmt.Println("Listening on port 8080")
  router.Run(":8080")
}

