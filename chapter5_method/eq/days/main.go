package main

import "fmt"

var s = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

type Day int

func (d Day) String() string {
	return s[d]
}

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	var day Day = 6
	fmt.Println(day)
}
