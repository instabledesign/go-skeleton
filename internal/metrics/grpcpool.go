package metrics

import (
	"time"

	"github.com/processout/grpc-go-pool"
)

type Pool struct {
	*grpcpool.Pool

	name     string
	recorder *Recorder

	done     chan struct{}
	interval *time.Ticker
}

func (p *Pool) GetName() string {
	return p.name
}

func (p *Pool) Close() {
	close(p.done)
	p.Pool.Close()
}

func (p *Pool) CollectMetrics() {
	p.done = make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-p.interval.C:
				p.recorder.CollectGrpcPoolState(p)
			case <-p.done:
				p.interval.Stop()
				return
			}
		}
	}()
}

func NewPool(name string, pool *grpcpool.Pool, rec *Recorder) *Pool {
	return &Pool{
		name:     name,
		Pool:     pool,
		recorder: rec,
	}
}
