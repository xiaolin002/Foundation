// E:/go_example/foundation/chapter5_method/eq/timezone/main.go
package main

import "fmt"

// TZ 为 int 定义别名类型
type TZ int

// 定义一些常量表示时区
const (
	UTC TZ = iota
	CST
	EST
	PST
)

// 定义一个 map，将时区的缩写映射为它的全称
var timezoneFullNames = map[TZ]string{
	UTC: "Coordinated Universal Time",
	CST: "China Standard Time",
	EST: "Eastern Standard Time",
	PST: "Pacific Standard Time",
}

// String 为类型 TZ 定义 String() 方法，它输出时区的全称
func (t TZ) String() string {
	return timezoneFullNames[t]
}

func main() {
	// 测试代码
	var currentTimezone TZ = 1
	fmt.Println(currentTimezone) // 输出: China Standard Time
}
