package fourq

import (
	"fmt"
)

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [2]uint64

func (e baseFieldElem) String() string {
	return fmt.Sprintf("%16.16x %16.16x", e[1], e[0])
}

//go:noescape
func bfeAdd(c, a, b *baseFieldElem)

//go:noescape
func bfeSub(c, a, b *baseFieldElem)

//go:noescape
func bfeMul(c, a, b *baseFieldElem)

//go:noescape
func bfeSquare(c, a *baseFieldElem)
