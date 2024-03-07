package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("/Users/hpfish/go/src/hutao/the-way-to-go/resource/goprogram")
	outputFile, _ := os.OpenFile("/Users/hpfish/go/src/hutao/the-way-to-go/resource/goprogramT",
		os.O_WRONLY|os.O_CREATE, 0644)
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	for {
		inputString, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err = writer.WriteString(outputString)
		if err != nil {
			return
		}
		writer.Flush()
	}

	fmt.Println("Conversion done")
}
