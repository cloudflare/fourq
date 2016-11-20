package fourq

import (
	"fmt"
)

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [2]uint64

func newBaseFieldElem() *baseFieldElem {
	return &baseFieldElem{}
}

func (e *baseFieldElem) String() string { return fmt.Sprintf("%16.16x %16.16x", e[1], e[0]) }

func (e *baseFieldElem) Set(a *baseFieldElem) { e[0], e[1] = a[0], a[1] }
func (e *baseFieldElem) SetZero()             { e[0], e[1] = 0, 0 }
func (e *baseFieldElem) SetOne()              { e[0], e[1] = 1, 0 }

func (e *baseFieldElem) IsZero() bool { return e[0] == 0 && e[1] == 0 }

func (e *baseFieldElem) Neg(a *baseFieldElem) *baseFieldElem {
	e[0] = ^a[0]
	e[1] = (^a[1]) & aMask
	return e
}

// chain1251 sets e to a^(2^125-1) and returns e.
func (e *baseFieldElem) chain1251(a *baseFieldElem) *baseFieldElem {
	t1 := newBaseFieldElem()
	t2 := newBaseFieldElem()
	t3 := newBaseFieldElem()
	t4 := newBaseFieldElem()
	t5 := newBaseFieldElem()

	bfeSquare(t2, a)
	bfeMul(t2, t2, a)
	bfeSquare(t3, t2)
	bfeSquare(t3, t3)
	bfeMul(t3, t3, t2)
	bfeSquare(t4, t3)
	bfeSquare(t4, t4)
	bfeSquare(t4, t4)
	bfeSquare(t4, t4)
	bfeMul(t4, t4, t3)
	bfeSquare(t5, t4)
	for i := 0; i < 7; i++ {
		bfeSquare(t5, t5)
	}
	bfeMul(t5, t5, t4)
	bfeSquare(t2, t5)
	for i := 0; i < 15; i++ {
		bfeSquare(t2, t2)
	}
	bfeMul(t2, t2, t5)
	bfeSquare(t1, t2)
	for i := 0; i < 31; i++ {
		bfeSquare(t1, t1)
	}
	bfeMul(t1, t1, t2)
	for i := 0; i < 32; i++ {
		bfeSquare(t1, t1)
	}
	bfeMul(t1, t2, t1)
	for i := 0; i < 16; i++ {
		bfeSquare(t1, t1)
	}
	bfeMul(t1, t1, t5)
	for i := 0; i < 8; i++ {
		bfeSquare(t1, t1)
	}
	bfeMul(t1, t1, t4)
	for i := 0; i < 4; i++ {
		bfeSquare(t1, t1)
	}
	bfeMul(t1, t1, t3)
	bfeSquare(t1, t1)
	bfeMul(e, t1, a)

	return e
}

// Invert sets e to a^(-1) and returns e.
func (e *baseFieldElem) Invert(a *baseFieldElem) *baseFieldElem {
	t := newBaseFieldElem().chain1251(a)
	bfeSquare(t, t)
	bfeSquare(t, t)
	bfeMul(e, t, a)

	return e
}

// reduce sets e to zero if it is equal to p. This is the only case where e will
// not naturally be reduce to canonical form.
func (e *baseFieldElem) reduce() {
	if e[0] == bMask && e[1] == aMask {
		e[0], e[1] = 0, 0
	}
}

//go:noescape
func bfeDbl(c, a *baseFieldElem)

func bfeHalf(c, a *baseFieldElem) {
	elem := &baseFieldElem{0x0000000000000000, 0x4000000000000000}
	bfeMul(c, elem, a)
}

//go:noescape
func bfeAdd(c, a, b *baseFieldElem)

//go:noescape
func bfeSub(c, a, b *baseFieldElem)

//go:noescape
func bfeMul(c, a, b *baseFieldElem)

//go:noescape
func bfeSquare(c, a *baseFieldElem)
