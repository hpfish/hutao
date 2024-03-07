package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	file, _ := os.OpenFile("/Users/hpfish/go/src/hutao/the-way-to-go/resource/vard.json",
		os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		panic(enc)
	}
	decoder := json.NewDecoder(file)
	var v VCard
	err = decoder.Decode(&v)
	if err != nil {
		return
	}
	fmt.Println(v)

}
