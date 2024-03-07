package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input2      string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input2, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("The input was : %s\n", input2)
	}
}
