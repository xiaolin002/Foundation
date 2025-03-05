package main

import (
	"context"
	"fmt"
	"time"
)

func doWorkWithContext(ctx context.Context) string {
	select {
	case <-time.After(2 * time.Second):
		return "Work done"
	case <-ctx.Done():
		return ctx.Err().Error()
	}
}

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result := doWorkWithContext(ctx)
	fmt.Println(result)
}
