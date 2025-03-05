package highFunc

/**
 * @Description
 * @Date 2025/2/28 22:55
 **/
type Car struct {
}

type Cars []Car

func (cs Cars) Process(f func(car Car)) {
	for _, c := range cs {
		f(c)
	}

}
