package main

/**
 * @Description
 * @Date 2025/1/16 15:52
 **/
//写一个函数将从start 到end索引的元素从切片删除
func delete(slice []int, start, end int) []int {

	if start < 0 || start >= len(slice) || end < 0 || end >= len(slice) || start > end {
		return slice
	}
	slice = slice[:start]
	slice = append(slice, slice[end+1:]...)
	return slice

}
