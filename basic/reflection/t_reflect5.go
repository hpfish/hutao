package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	Username string
	Password string
}

func main() {
	v := reflect.ValueOf(&Account{})
	fmt.Println(v.Type().Name()) // ""
	indirectV := reflect.Indirect(v)
	fmt.Println(indirectV.Type().Name()) // Account
}
