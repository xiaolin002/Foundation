package main

import "fmt"

type Any any

type Car struct {
	Model        string // 型号
	Manufacturer string // 制造商
	BuildYear    int    // 建造年份
	//...
}
type Cars []*Car

func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		//f(c) 调用传入的函数 f，
		//并将当前的 *Car 指针 c 作为参数传递给它
		f(c)
	}
}

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) any) []any {
	result := make([]any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result

}

func MakeSortedAppender(manufactures []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufactures {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	// 定义一个匿名函数，用于将汽车添加到 sortedCars 中
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars

}

func main() {
	ford := &Car{"Fiesta", "Ford", 2010}
	bmw := &Car{"XL 450", "BMW", 2016}
	merc := &Car{"D600", "Mercedes", 2019}
	bmw2 := &Car{"X 800", "BMW", 2018}
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(c *Car) bool {
		return (c.Manufacturer == "BMW") && (c.BuildYear > 2010)
	})
	// 修改部分
	fmt.Println("AllCars: ")
	for _, car := range allCars {
		fmt.Println(car)
	}
	fmt.Println("New BMWs: ")
	for _, car := range allNewBMWs {
		fmt.Println(car)
	}
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWs := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWs, " BMWs")

}
