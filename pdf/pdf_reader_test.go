package pdf

import (
	"os"
	"testing"
)

func TestPdfReader(t *testing.T) {
	pdfFile, err := os.Open(os.Getenv("PWD") + "/test.pdf")
	if err != nil {
		println("Error open test pdf file ", err)
		return
	}
	for true {
		partOfFile := make([]byte, 100)
		read, err := pdfFile.Read(partOfFile)
		if read == 0 {
			println("Found end of file")
			return
		}
		if err != nil {
			println("Error read part of file")
			return
		}
		println("Part of file: ", string(partOfFile))
	}
}
