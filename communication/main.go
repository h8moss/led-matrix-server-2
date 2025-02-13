package communication

import (
	"fmt"
	"os"
	"syscall"
)

func CreatePipe(path string) {
	os.Remove(path)
	err := syscall.Mkfifo(path, 0644)
	if err != nil {
		panic(err)
	}
}

func InitializePipe(path string, c chan string) {
	f, err := syscall.Open(path, syscall.O_CREAT|syscall.O_WRONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		panic(err)
	}

	for command := range c {
		fmt.Printf("Received command: %s\n", command)
		buf := []byte(command)
		syscall.Write(f, buf)
	}

	syscall.Close(f)
}
