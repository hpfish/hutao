package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"unsafe"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

type Header struct{}

type Codec interface {
	ReadHeader(*Header) error
}

var _ Codec = (*GobCodec)(nil)

func (g *GobCodec) ReadHeader(*Header) error {
	return nil
}

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {

	fmt.Println(unsafe.Alignof(demo1{}))
	fmt.Println(unsafe.Alignof(demo2{}))
	fmt.Println(unsafe.Sizeof(demo1{}))
	fmt.Println(unsafe.Sizeof(demo2{}))
	fmt.Println(unsafe.Sizeof(demo3{}))
	fmt.Println(unsafe.Sizeof(demo4{}))

}
