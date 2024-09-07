package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"syscall"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Config struct {
        LedMatrixPath string
        FifoPath      string
        Port   			  int
}

func main() {
	fmt.Println("Reading config...")
	content, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Unpacking config...")
	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("--Config found! Now starting server...")

	fmt.Println("Creating communication pipe...")
	os.Remove(config.FifoPath)
	err = syscall.Mknod(config.FifoPath, syscall.S_IFIFO|0666, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting led matrix manager")
  matrixManagerPath := config.LedMatrixPath + "/led-matrix-manager"
  cmd := exec.Command("sudo", matrixManagerPath, config.FifoPath)
  err = cmd.Start()
  if err != nil {
    panic(err)
  }

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/index.html", false)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	fmt.Println("Listening on port 8080")
	router.Run(":8080")
}

