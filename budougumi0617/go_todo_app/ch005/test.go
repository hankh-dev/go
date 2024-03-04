package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("중단되지 않음")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)

	fmt.Println("hm?")
}
