// Package utilities does what you think it does
package utilities

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add does what you think it does
// [Adding docs]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](i1 T, i2 T) T {
	return i1 + i2
}
