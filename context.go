package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("time line exceeded of 2 seconds")
	}
	time.Sleep(time.Second * 3)
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("Doing something bakwaas")
		}
	}
}

/*
	func main() {
		ctx := context.Background()
		seedContext(ctx)
		readCtx(ctx)
	}
*/
func readCtx(ctx context.Context) {
	value := ctx.Value("one")
	fmt.Println(value)
}
func seedContext(ctx context.Context) {
	ctx = context.WithValue(ctx, "key", "value")
	return

}
