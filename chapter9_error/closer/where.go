package main

import (
	"log"
	"runtime"
)

/**
 * @Description
 * @Date 2025/2/28 22:52
 **/

func main() {

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Println("%s:%d", file, line)
	}

	where()
	where()
	where()

}
