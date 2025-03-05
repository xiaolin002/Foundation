package main

import "fmt"

// 判断是不是素数
func main() {

	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime, " ")
		out := make(chan int)
		// 过滤器
		go filter(ch, out, prime)
		ch = out
	}

}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}
func filter(in chan int, out chan int, prime int) {
	for {

		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

/*
generate 一直在生成数 当2 进入for循环中 打印出来 此时ch为空的，但是ch=out 比go filter 执行快
，所以ch此时为空，进入下一次for循环，ch为空 所以一直阻塞在这里prim:=<-ch 所以只能等filter协程，
filter第一次进入时ch为空 ，需要等generate生成数据，故i=3，进入out 此时主程序中for循环执行，
再次启动一个协程filter ，上一个协程filter仍然在接受generate的数据 直到数据完成，
相当与filter往out中传多少数就会执行多少for循
*/
