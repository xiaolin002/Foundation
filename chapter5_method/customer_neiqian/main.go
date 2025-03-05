package main

import "fmt"

type Log struct {
	msg string
}

// Customer 内嵌类型时出现的方法
type Customer struct {
	Name string
	Log
}

func main() {
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2- after me the world will be a better place")
	fmt.Println(c)
	///调用 fmt.Println(c)，触发 Customer 的 String 方法。
	//在 Customer 的 String 方法中，调用 Log 的 String 方法

}

func (l *Log) Add(s string) {
	l.msg += "\n" + s

}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprint(c.Log)
}
