package main

import "fmt"

const (
	WIDTH  = 10
	HEIGHT = 4
)

type pixel int

// 声明一个数组 但没有初始化
var screen [WIDTH][HEIGHT]pixel

func main() {
	fmt.Println(screen)

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = pixel(x + y)
			fmt.Print(screen[x][y])

		}
		fmt.Println()

	}

}
