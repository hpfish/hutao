package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/hpfish/go/src/hutao/the-way-to-go/resource/produccts.txt")
	if err != nil {
		return
	}
	reader := csv.NewReader(file)

	for {
		lines, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
