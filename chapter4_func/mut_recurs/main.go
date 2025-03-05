package main

import "fmt"

func main() {

	fmt.Printf("even(16) = %v\n", even(16))
	fmt.Printf("odd(17) = %v\n", odd(17))
	fmt.Printf("even(18) = %v\n", even(18))

}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)

}

func odd(nr int) bool {

	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
