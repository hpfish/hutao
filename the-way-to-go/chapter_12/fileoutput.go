package main

import (
	"bufio"
	"os"
)

func main() {
	outputFile, err := os.OpenFile("/Users/hpfish/go/src/hutao/the-way-to-go/resource/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	outputStr := "hello world!\n"

	for i := 0; i < 10; i++ {
		writer.WriteString(outputStr)
	}
	writer.Flush()
}
