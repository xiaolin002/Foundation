package main

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/1 11:40
 **/
func TestFlag(t *testing.T) {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", "--", "separator")
	flag.Parse()
	fmt.Print(strings.Join([]string{"你好", "xiaoming"}, *sep))
	if !*n {
		fmt.Println()
		fmt.Println("middle----")
	}
	fmt.Println("end----")
}
