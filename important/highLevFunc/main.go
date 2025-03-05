package main

// 高阶函数

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

func main() {

}
