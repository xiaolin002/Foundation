package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) First3Chars() string {
	return t.Time.String()[0:4]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("Full time now", m.String())
	fmt.Println("First 3 chars", m.First3Chars())
}
