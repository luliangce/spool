package main

import (
	"context"
	"time"

	"github.com/luliangce/spool"
)

func main() {
	pool := spool.NewPool(10) //define a pool size

	for i := 0; i < 100; i++ {
		pool.Go(context.Background(), func() {
			time.Sleep(time.Millisecond * 100)
		})
	}
	pool.Wait()
}
