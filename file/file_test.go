package file

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileGo(t *testing.T) {
	CreateRandomFile()

	assert.Equal(t, -1, -1)
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
