package main

import (
	"os"
	"strconv"
)

// 向控制台输出的话使用os.Stdout.WriteString()

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func print(any ...interface{}) {
	for i, arg := range any {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch v := arg.(type) {
		case Stringer:
			os.Stdout.WriteString(v.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(v))
		case string:
			os.Stdout.WriteString(v)
		default:
			os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "hello", Celsius(100.0))
}
