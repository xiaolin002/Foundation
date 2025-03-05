package main

/**
 * @Description
 * @Date 2025/1/16 16:00
 **/
// 将元素添加到旧元素后面设计到扩容
func replace(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := len(data) + m
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice

}
