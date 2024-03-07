package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

type VCard2 struct {
	FirstName string
	LastName  string
	Addresses []*Address2
	Remark    string
}

var content string

func main() {
	pa := &Address2{"private", "Aartselaar", "Belgium"}
	wa := &Address2{"work", "Boom", "Belgium"}
	vc := VCard2{"Jan", "Kersschot", []*Address2{pa, wa}, "none"}
	file, _ := os.OpenFile("/Users/hpfish/go/src/hutao/the-way-to-go/resource/vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}

	dec := gob.NewDecoder(file)
	var v VCard2
	err = dec.Decode(&v)
	if err != nil {
		return
	}
	fmt.Println(v)

}
