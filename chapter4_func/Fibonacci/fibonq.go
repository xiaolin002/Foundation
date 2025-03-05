package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		result := fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//  改造后
//func main() {
//	eleven_1 i := 0; i < 10; i++ {
//		position, value := fibonacci(i)
//		fmt.Printf("fibonacci(%d) is: %d (position: %d)\n", i, value, position)
//	}
//}
//
//func fibonacci(n int) (position int, res int) {
//	position = n
//	if n <= 1 {
//		res = n // 0 eleven_1 fibonacci(0) and 1 eleven_1 fibonacci(1)
//	} else {
//		_, res = fibonacci(n - 1)
//		_, res2 := fibonacci(n - 2)
//		res += res2
//	}
//	return
//}
