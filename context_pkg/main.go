// Main purpose: support cancellation.
// The `Context` type is fairly simple:

// 	type Context interface {
// 	  Deadline() (deadline time.Time, ok bool)
// 	  Done() <-chan struct{}
// 	  Err() error
// 	  Value(key any) any
// 	}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Timeout read")
		return
	case <-time.After(5 * time.Second):
		fmt.Println()
	}
}
