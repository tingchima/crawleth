// Package dao provides
package dao

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"

	"portto_interview/indexer-service/pkg/convert"

	"github.com/go-redis/redis/v8"
)

// SetBlockNumber .
func (d *Dao) SetBlockNumber(ctx context.Context, key string, val *big.Int) error {
	// Zero expiration means the key has no expiration time.
	b, _ := json.Marshal(val)
	return d.Cache().Set(ctx, key, b, 0).Err()
}

// GetBlockNumber .
func (d *Dao) GetBlockNumber(ctx context.Context, key string) (*big.Int, error) {
	val, err := d.Cache().Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return convert.StrTo(val).BigInt(), nil
}

// AddBlockNumbers .
func (d *Dao) AddBlockNumbers(ctx context.Context, key string, numbers ...*big.Int) error {
	values := make([]interface{}, len(numbers))
	for i := range numbers {
		b, _ := json.Marshal(numbers[i])
		values[i] = b
	}
	return d.Cache().SAdd(ctx, key, values...).Err()
}

// GetBlockNumbers .
func (d *Dao) GetBlockNumbers(ctx context.Context, key string) ([]*big.Int, error) {
	members, err := d.Cache().SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	numbers := make([]*big.Int, len(members))
	for i := range members {
		number, _ := strconv.ParseInt(members[i], 10, 64)
		numbers[i] = big.NewInt(number)
	}
	return nil, nil
}

// PopBlockNumber .
func (d *Dao) PopBlockNumber(ctx context.Context, key string) (*big.Int, error) {
	val, err := d.Cache().SPop(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return convert.StrTo(val).BigInt(), nil
}

// PopBlockNumbers .
func (d *Dao) PopBlockNumbers(ctx context.Context, key string, cnt int64) ([]*big.Int, error) {
	values, err := d.Cache().SPopN(ctx, key, cnt).Result()
	if err != nil {
		return nil, err
	}
	numbers := make([]*big.Int, len(values))
	for i := range values {
		numbers[i] = convert.StrTo(values[i]).BigInt()
	}
	return numbers, nil
}
