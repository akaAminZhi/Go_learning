package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("context cancel")
	}
}

func main() {
	//这是一个空的context
	ctx := context.Background()
	//接着使用cancel
	ctx, cancel := context.WithCancel(ctx)

	cancel()

}
