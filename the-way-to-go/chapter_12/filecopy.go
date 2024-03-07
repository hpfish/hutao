package main

import (
	"io"
	"os"
)

func main() {
	CopyFile("/Users/hpfish/go/src/hutao/the-way-to-go/resource/target.txt", "/Users/hpfish/go/src/hutao/the-way-to-go/resource/source.txt")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
