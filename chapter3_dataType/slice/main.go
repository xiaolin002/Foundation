package main

import "fmt"

func main() {

	//stack = append(stack, v)     压栈
	//top := stack[len(stack)-1]   stack的顶部元素
	//stack = stack[:len(stack)-1]  通过收缩来弹出栈顶的元素

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s, 2))

}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func remove2(slice []int, i int) []int {
	// 这里的slice[i] 直接被最后一个元素替换了
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
