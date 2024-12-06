package customScanner

import (
	"bufio"
	"io"
	"log"
	"os"
)

type CustomScanner struct {
	io.Closer
	*bufio.Scanner
}

func GetFileScanner(fileName string) *CustomScanner {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return &CustomScanner{file, scanner}
}
