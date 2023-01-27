package file

import (
	"math/rand"
	"os"
	"strconv"
)

func CreateRandomFile() {
	var pwd = os.Getenv("PWD")
	println(pwd)
	var file, _ = os.Create(pwd + "/test.txt")
	for i := 0; i < 1024; i++ {
		var word = []byte(strconv.Itoa(rand.Int()))
		file.Write(word)
		file.WriteString("\n")
	}
}
