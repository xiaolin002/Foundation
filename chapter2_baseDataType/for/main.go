package main

import "fmt"

func main() {
	values := []int{1, 2, 3}
	var ptrs []*int

	for _, val := range values {
		ptrs = append(ptrs, &val)
	}

	for _, ptr := range ptrs {
		fmt.Println(*ptr)
	}

	// 这里的输出结果是3 3 3 因为val是一个局部变量，每次循环都会被重新赋值
	// 所以ptr指向的是同一个val，所以每次输出的都是最后一个val的值

	/*最后在1.22 版本中，Go 语言的编译器会对这个问题进行优化，
	将val的值复制到一个新的变量中，然后将这个变量的地址赋给ptr，
	这样就可以避免这个问题了。
	*/
	/*

	 values := []int{1, 2, 3}
	    var ptrs []*int

	    for _, val := range values {
	        ptrs = append(ptrs, &val)
	    }

	    for _, ptr := range ptrs {
	        fmt.Println(*ptr)
	    }
	*/

}
