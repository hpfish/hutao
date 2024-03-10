package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//func (c Cat) Quack() {
//	fmt.Println("meow")
//}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

//func main() {
//	//var c Duck = Cat{}
//	//c.Quack()
//}
