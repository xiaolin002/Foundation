package main

import "fmt"

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥", 4: "你好"}

	fmt.Println(RMB, symbol[4]) //

}
