package test

import (
	"fmt"
	"testing"
)

/**
 * @Description
 * @Date 2025/1/15 10:26
 **/

func TestPrintSlice(t *testing.T) {

	a := []string{"a", "b", "c", "d"}
	for i, _ := range a {
		fmt.Println(i, "is", a[i])
	}

}
