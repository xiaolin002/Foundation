package main

import (
	"fmt"
	"time"
)

/*
  写一个通道证明它的阻塞性 开启一个协程接收通道的数据 持续15秒
  然后给通道放入一个值 在不同的阶段打印消息并观察输出
*/

func main() {

	ch := make(chan string)
	st := time.Now()
	go func() {
		// 用来接收通道的数据
		startTime := time.Now()
		//时间判断必须放在for循环中
		// 否则只会判断一次
		for {
			if time.Since(startTime) >= 15*time.Second {
				fmt.Println("通道已经持续15秒")
				break
			}
			data, ok := <-ch
			if ok {
				fmt.Printf("协程接收到通道数据: %s\n", data)
			}
		}

	}()

	time.Sleep(5 * time.Second)
	ch <- "hello"
	time.Sleep(5 * time.Second)
	ch <- "world"
	time.Sleep(5 * time.Second)
	// 关闭通道
	close(ch)
	fmt.Println(time.Since(st).Seconds())
	fmt.Println("程序结束")

}

/*
for {
			if time.Since(startTime) > 15*time.Second {
				fmt.Println("通道已经持续15秒")
				return
			}
			select {
			case data, ok := <-ch:
				if ok {
					fmt.Printf("协程接收到通道数据: %s\n", data)
				} else {
					fmt.Println("通道已关闭")
					return
				}
			default:
				// 模拟一些工作
				time.Sleep(100 * time.Millisecond)
			}
		}


*/
