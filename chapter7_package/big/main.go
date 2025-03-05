package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	//((math.MaxInt64 * math.MaxInt64) + math.MaxInt64) / 1956
	fmt.Printf("Big Int: %v\n", ip)

	//
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)

	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)

}
