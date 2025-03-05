package main

import "fmt"

// Reader 定义一个简单的读取接口
type Reader interface {
	Read() string
}

// Writer 定义一个简单的写入接口
type Writer interface {
	Write(data string)
}

// ReadWriter 嵌套 Reader 和 Writer 接口
type ReadWriter interface {
	Reader
	Writer
}

// File 结构体实现了 ReadWriter 接口
type File struct{}

func (f File) Read() string {
	return "读取文件内容"
}

func (f File) Write(data string) {
	fmt.Printf("写入文件: %s\n", data)
}

func main() {
	var rw ReadWriter = File{}
	var r Reader = rw
	var w Writer = rw

	fmt.Println(r.Read())
	w.Write("新的数据")
}
