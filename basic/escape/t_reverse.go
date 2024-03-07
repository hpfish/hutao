package main

import (
	"io"
	"os"
)

func main() {
	//Println(string(*ReverseA("Ding Ding"))) // ???
	result := []rune("Ding Ding")
	ReverseB(result)
	Println(string(result))
}

func Println(str string) {
	io.WriteString(os.Stdout,
		str+"\n")
}

func ReverseB(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func ReverseA(str string) *[]rune {
	result := make([]rune, 0, len(str))
	for _, v := range []rune(str) {
		v := v
		defer func() {
			result = append(result, v)
		}()
	}
	return &result
}
