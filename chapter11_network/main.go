package main

import "net"

func main() {
	_, err := net.Dial("", "")
	if err != nil {
		println(err)
	}

}
