package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("02 Jan 2006 15:04"))

	fmt.Println(t.UTC())
}
