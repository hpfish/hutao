package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string // 标题
	Body  []byte
}

func save(p *Page) error {
	err := ioutil.WriteFile(p.Title, p.Body, 0666)
	if err != nil {
		return err
	}
	return nil
}

func load(title string) *Page {
	file, err := ioutil.ReadFile(title)
	if err != nil {
		return nil
	}
	return &Page{title, file}
}

func main() {
	p := Page{
		Title: "text.txt",
		Body:  []byte("hello world"),
	}
	err := save(&p)
	if err != nil {
		return
	}

	loadP := load("text.txt")
	fmt.Println(string(loadP.Body))
}
