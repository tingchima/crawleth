// Package convert provides
package convert

import (
	"math/big"
	"strconv"
)

// StrTo .
type StrTo string

// String .
func (st StrTo) String() string {
	return string(st)
}

// Int .
func (st StrTo) Int() (int, error) {
	return strconv.Atoi(st.String())
}

// MustInt .
func (st StrTo) MustInt() int {
	val, _ := st.Int()
	return val
}

// MustInt64 .
func (st StrTo) MustInt64() int64 {
	val, _ := strconv.ParseInt(st.String(), 10, 64)
	return val
}

// UInt64 .
func (st StrTo) UInt64() (int32, error) {
	val, err := strconv.Atoi(st.String())
	return int32(val), err
}

// MustUInt64 .
func (st StrTo) MustUInt64() int32 {
	val, _ := st.UInt64()
	return val
}

// BigInt .
func (st StrTo) BigInt() *big.Int {
	return big.NewInt(st.MustInt64())
}
