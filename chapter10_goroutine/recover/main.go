package main

func main() {

}

type Work struct {
	ID   int
	Data []byte
}

func server(workChan <-chan *Work) {
	for work := range workChan {

		go safelyDo(work)

	}

}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			// 处理错误
		}
	}()
	//do(work)
}
