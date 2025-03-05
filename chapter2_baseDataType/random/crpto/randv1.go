package crpto

import (
	"crypto/rand"
	"fmt"
)

/**
 * @Description
 * @Date 2025/1/11 16:06
 **/

func Rand100() {
	rBytes := make([]byte, 16)
	_, err := rand.Read([]byte(rBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", rBytes)

}
