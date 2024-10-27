package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()

	go func() { fmt.Println("別のゴルーチン") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("そして時は動き出す")
}
