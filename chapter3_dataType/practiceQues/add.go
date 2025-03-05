package main

/**
 * @Description
 * @Date 2025/1/16 15:52
 **/

//给定一个int类型切片 s和一个int类型的因子 扩展s使其长度为len（s）*factor

func add(s []int, factor int) []int {
	newLen := len(s) * factor
	newSlice := make([]int, newLen)
	copy(newSlice, s)
	return newSlice

}
