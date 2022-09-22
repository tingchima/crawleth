// Package app provides
package app

import "fmt"

// Order .
type Order string

const (
	// Desc .
	Desc Order = "desc"
	// Asc .
	Asc Order = "asc"
)

// Sorting .
type Sorting struct {
	Field string
	Order Order
}

// SortBy .
func (s Sorting) SortBy() string {
	if s.Order == "" {
		s.Order = Asc
	}
	return fmt.Sprintf("%s %s", s.Field, s.Order)
}
