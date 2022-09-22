// Package convert provides
package convert

// IntTo .
type IntTo int

// Int .
func (i IntTo) Int() int {
	return int(i)
}

// Bool .
func (i IntTo) Bool() bool {
	return i == 1
}
