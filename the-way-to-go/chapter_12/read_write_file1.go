package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "/Users/hpfish/go/src/hutao/the-way-to-go/chapter_12/input.dat"
	outputFile := "/Users/hpfish/go/src/hutao/the-way-to-go/chapter_12/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
