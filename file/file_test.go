package file

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileGo(t *testing.T) {
	var pwd = os.Getenv("PWD")
	CreateRandomFile(pwd + "test.txt")

	assert.Equal(t, -1, -1)
}

func TestSocketFileSender(t *testing.T) {
	var pwd = os.Getenv("PWD")
	// Init send folder
	os.Mkdir(pwd+"/sender", 0777)
	// Receive send folder
	os.Mkdir(pwd+"/receiver", 0777)
	// Init file
	file := CreateRandomFile(pwd + "/sender/test.txt")

	go func() string {
		fmt.Println("Try to start server")
		listen, err := net.Listen("tcp", ":8013")
		if err != nil {
			fmt.Println("Error listen to tcp 8013")
		}
		serverConnection, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("Accept listen to tcp 8013")
		}
		file, _ := os.Create(pwd + "/receiver/test.txt")
		for {
			// get message, output
			message, _ := bufio.NewReader(serverConnection).ReadBytes('\n')
			fmt.Print("Message Received:", string(message))
			_, err3 := file.Write(message)
			if err3 != nil {
				fmt.Println("Error write to file ", err3)
			}
		}
	}()

	go func() {
		fmt.Println("Try to start client")
		dial, err := net.Dial("tcp", "localhost:8013")
		if err != nil {
			fmt.Println("Error create client tcp to 8013")
		}
		var offset = int64(0)
		for i := 0; i < 10; i++ {
			file.Seek(offset, 0)
			partOfFile := make([]byte, 20)
			file.Read(partOfFile)
			_, err := dial.Write(partOfFile)
			if err != nil {
				fmt.Println("Error write to server ", err)
			}
			offset += 20
			time.Sleep(50 * time.Millisecond)
		}
	}()

	time.Sleep(60 * time.Second)
}

func TestSocket(t *testing.T) {

	go func() {
		fmt.Println("Try to start server")
		listen, err := net.Listen("tcp", ":8013")
		if err != nil {
			fmt.Println("Error listen to tcp 8013")
		}
		serverConnection, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("Accept listen to tcp 8013")
		}
		for {
			// get message, output
			message, _ := bufio.NewReader(serverConnection).ReadString('\n')
			fmt.Print("Message Received:", string(message))
		}
	}()

	go func() {
		fmt.Println("Try to start client")
		dial, err := net.Dial("tcp", "localhost:8013")
		if err != nil {
			fmt.Println("Error create client tcp to 8013")
		}
		for true {
			_, err := dial.Write([]byte("Hello to server!\n"))
			if err != nil {
				fmt.Println("Error sent message client tcp to 8013")
			}
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(60 * time.Second)
}
