// Package subscriber provides
package subscriber

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

// subEthHeader .
type subEthHeader struct {
	headers  chan *big.Int
	errs     chan error
	duration time.Duration
}

// NewEthHeader .
func NewEthHeader(duration time.Duration) Subscriber {
	return &subEthHeader{
		headers:  make(chan *big.Int),
		errs:     make(chan error),
		duration: duration,
	}
}

// Subscribe .
func (e *subEthHeader) Subscribe(ctx context.Context, f subscribeFunc) {
	fmt.Printf("start to subscribe new header...\n")
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(e.errs)
				close(e.headers)

			default:
				time.Sleep(e.duration * time.Second)
				headerNumber, err := f(ctx)
				if err != nil {
					e.errs <- err
				}
				e.headers <- headerNumber
			}
		}
	}()
}

func (e *subEthHeader) Err() <-chan error {
	return e.errs
}

func (e *subEthHeader) Headers() <-chan *big.Int {
	return e.headers
}
