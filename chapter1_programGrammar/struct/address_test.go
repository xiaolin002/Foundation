package main

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/1 11:41
 **/

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}
func TestAddress(t *testing.T) {

	v := 1
	incr(&v) // side effect: v is now 2
	fmt.Println(incr(&v))

}
