package spool

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Pool struct {
	weight *semaphore.Weighted
	wg     *sync.WaitGroup
}

func (p *Pool) Go(ctx context.Context, task func()) {
	p.wg.Add(1)
	go func() {
		err := p.weight.Acquire(ctx, 1)
		defer p.wg.Done()
		if err != nil {
			return
		}

		defer p.weight.Release(1)
		task()
	}()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func NewPool(size int64) *Pool {
	return &Pool{
		weight: semaphore.NewWeighted(size),
		wg:     new(sync.WaitGroup),
	}
}
