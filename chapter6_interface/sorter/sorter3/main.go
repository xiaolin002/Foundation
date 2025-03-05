package main

import (
	"fmt"
	"sort"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntSlice(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fatal error")
	}
	fmt.Printf("The sorted array is:%v\n", a)

}

func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "saturday"}
	a := sort.StringSlice(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fatal error")
	}
	fmt.Printf("The sorted array is:%v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}
func (p *dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}
func (p *dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}
func days() {
	Sunday := day{0, "sun", "Sunday"}
	Monday := day{1, "mon", "Monday"}
	Tuesday := day{2, "tue", "Tuesday"}
	Wednesday := day{3, "wed", "Wednesday"}
	Thursday := day{4, "thu", "Thursday"}
	Friday := day{5, "fri", "Friday"}
	Saturday := day{6, "sat", "Saturday"}
	data := []*day{&Sunday, &Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fatal error")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}
func main() {
	ints()
	strings()
	days()

}
