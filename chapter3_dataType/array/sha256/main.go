package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
)

var hashType string
var inputString string

func main() {

	flag.StringVar(&hashType, "hash", "sha256", "哈希算法类型，可选值：sha256, sha384, sha512")
	flag.StringVar(&inputString, "input", "hello", "输入字符串")
	flag.Parse()

	var h hash.Hash
	switch hashType {
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		println("哈希类型错误")
		return
	}

	_, err := io.WriteString(h, inputString)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(h.Sum([]byte{}))

}
