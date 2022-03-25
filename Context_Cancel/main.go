package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Finish do something")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("context cancel yo")
	}
}

func main() {
	//这是一个空的context
	ctx := context.Background()
	//接着使用cancel
	ctx, cancel := context.WithCancel(ctx)
	// 创建一个进程去执行 cancel，2秒后
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomething(ctx)

}
