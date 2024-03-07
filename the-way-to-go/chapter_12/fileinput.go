package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("/Users/hpfish/go/src/hutao/the-way-to-go/chapter_12/input.dat")
	if inputError != nil {
		panic(inputError)
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		//inputString, readerError := inputReader.ReadString('\n')
		//fmt.Printf("The input was : %s", inputString)
		//if readerError == io.EOF {
		//	return
		//}

		buf := make([]byte, 1024)
		n, err := inputReader.Read(buf)
		if err == io.EOF {
			fmt.Println(string(buf))
			return
		}
		if n == 0 {
			fmt.Println(string(buf))
			break
		}

	}
}
