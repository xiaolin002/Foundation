package _type

/**
 * @Description
 * @Date 2025/1/1 18:59
 **/

type FT func(a string, b int) (ok bool, err error)

// 别名  但是上边的方法可以有其他得理解
// type FT = func(a string, b int) (bool, err error)

func MyFunction(a string, b int) (bool, error) {
	// 在这里实现函数的具体逻辑
	// ...
	return true, nil // 或者根据实际情况返回其他值
}

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}
