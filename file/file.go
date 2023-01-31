package file

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func CreateRandomFile() {
	var pwd = os.Getenv("PWD")
	println(pwd)
	var file, _ = os.Create(pwd + "/test.txt")
	for i := 0; i < 10; i++ {
		var word = []byte(strconv.Itoa(rand.Int()))
		file.Write(word)
		file.WriteString("\n")
	}
	var offset = int64(0)
	for i := 0; i < 10; i++ {
		file.Seek(offset, 0)
		b1 := make([]byte, 20)
		n1, _ := file.Read(b1)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
		offset += 20
	}

}
