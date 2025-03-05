package main

/**
 * @Description
 * @Date 2025/1/16 15:52
 **/
//写一个函数 将切片插入到另一个切片的制定为止
func insert(slice, insertion []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	makeSlice := make([]int, len(slice)+len(insertion))
	// 将索引之前的数据复制到新切片
	copy(makeSlice, slice[:index])
	// 将插入的数据复制到新切片之后
	copy(makeSlice[index:], insertion)
	// 将索引之后的数据复制到新切片之后
	copy(makeSlice[index+len(insertion):], slice[index:])
	return makeSlice

}
