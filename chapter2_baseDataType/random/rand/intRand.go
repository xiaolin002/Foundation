package rand100

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Description
 * @Date 2025/1/11 12:04
 **/

func IntRand100() {

	//rand.Seed(time.Now().UnixMilli())
	//randomNumber := rand.Intn(100) + 1
	//fmt.Println(randomNumber)

	// 从1.20版本后，math/rand包不再使用全局的随机数生成器，而是使用一个本地的随机数生成器。

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		rand11 := rand.Intn(100) + 1
		fmt.Println(rand11)
	}

	rand.Int()

}
