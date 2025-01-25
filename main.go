package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"example.com/api"
	"example.com/communication"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Config struct {
	LedMatrixPath string
	FifoPath      string
	Port          int
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

	fmt.Println("Config found! Now starting server...")

	fmt.Println("Creating communication pipe...")
	communication.CreatePipe(config.FifoPath)
	pipeChan := make(chan string, 100)

	fmt.Println("Starting led matrix manager")
	matrixManagerPath := config.LedMatrixPath + "/build/led-matrix-manager"
	cmd := exec.Command("sudo", matrixManagerPath, config.FifoPath)
	fmt.Println(cmd.Args)

	// Get the output pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	// Create a goroutine to handle stdout
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			// Print each line from stdout as it's received
			trimmed := strings.Trim(scanner.Text(), "\n\t ")
			if trimmed[len(trimmed)-1:] != "0" {
				fmt.Println("led-matrix-manager: ", scanner.Text())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading stdout:", err)
		}
	}()

	// Create a goroutine to handle stderr
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			// Print each line from stderr as it's received
			fmt.Println("stderr:", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading stderr:", err)
		}
	}()

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

	fmt.Println("Opening pipe...")
	go communication.InitializePipe(config.FifoPath, pipeChan)

	fmt.Println("Initializing server...")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	// Setup route group for the API
	api.CreateApi(router, pipeChan)
	fmt.Println("Listening on port " + strconv.Itoa(config.Port))
	router.Run(":" + strconv.Itoa(config.Port))
}
