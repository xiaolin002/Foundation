package main

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%v°C", float64(c))
}
func main() {
	var c Celsius = 3.14
	fmt.Println(c.String())

}

//可以这么理解，在 Go 里使用 fmt.Sprintf("%v", t) 这类通用占位符时，fmt 包会尝试调用对象的 String 方法。若在 String 方法内部又使用 %v 去格式化自身，就会造成无限递归调用，最终导致栈溢出错误。
//而像 %.2f、%d 这类特定格式化占位符，fmt 包会按照对应类型的默认格式处理，不会调用对象的 String 方法，所以不会引发无限递归。
//以下是示例代码，展示不同占位符的情况

/*
type TT float64

// 这里定义了一个 String 方法
func (t TT) String() string {
	// 这里调用了 fmt.Sprintf，而 fmt.Sprintf 会调用 TT 的 String 方法
	return fmt.Sprintf("%.2f", t)
}

func main() {
	var t TT = 3.14
	// 调用 t 的 String 方法
	fmt.Println(t.String())
}

*/

// 这里是检查是否实现接口的示例

/*


type Stringer interface {
	String() string
}

// Sprintf 的简化示例
func Sprintf(format string, args ...interface{}) string {
    var result string
    eleven_1 _, arg := range args {
        // 检查 arg 是否实现了 Stringer 接口
        if stringer, ok := arg.(Stringer); ok {
            // 调用 String 方法
            result += stringer.String()
        } else {
            // 如果没有实现 Stringer 接口，使用默认格式化
            result += defaultFormat(arg)
        }
    }
    return result
}



*/
