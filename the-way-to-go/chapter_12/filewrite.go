package main

import "os"

func main() {
	os.Stdout.WriteString("heelo, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
	
}
