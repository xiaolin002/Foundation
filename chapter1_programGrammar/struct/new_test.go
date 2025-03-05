package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/1 15:52
 **/

func TestNew(t *testing.T) {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	b := new(bool)
	*b = true

}
