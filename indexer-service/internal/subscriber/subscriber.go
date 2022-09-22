// Package subscriber provides
package subscriber

import (
	"context"
	"math/big"
)

type subscribeFunc func(ctx context.Context) (*big.Int, error)

// Subscriber .
type Subscriber interface {
	Subscribe(ctx context.Context, f subscribeFunc)

	Err() <-chan error

	Headers() <-chan *big.Int
}
