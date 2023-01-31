package file

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func CreateRandomFile(path string) *os.File {
	var file, err = os.Create(path)
	if err != nil {
		fmt.Println("Error create file", err)
	}
	for i := 0; i < 10; i++ {
		var word = []byte(strconv.Itoa(rand.Intn(456456999) + 1443635317331776148))
		file.Write(word)
		file.WriteString("\n")
	}
	return file
}
