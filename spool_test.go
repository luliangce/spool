package spool_test

import (
	"context"
	"testing"
	"time"

	"github.com/luliangce/spool"
)

func TestPool(t *testing.T) {
	pool := spool.NewPool(4)
	for i := 0; i < 8; i++ {
		pool.Go(context.Background(), func() {
			time.Sleep(time.Millisecond * 10)
		})
	}
	pool.Wait()
}
