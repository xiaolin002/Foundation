package rand8str

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Description
 * @Date 2025/1/11 16:33
 **/

func RandStrOrNum8() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//A  45  a   97
	//生城100个8位随机字符串
	str := make([]string, 100)

	// 生成1个8位随机字符串  []byte字节类型 0-255
	randStr := make([]byte, 8)
	for i := 0; i < 100; i++ {
		for j := 0; j < 8; j++ {
			randStr[j] = charset[r.Intn(len(charset))]
		}
		fmt.Println(string(randStr))
		str = append(str, string(randStr))
	}

}
