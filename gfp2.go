package fourq

import (
	"fmt"
	"math/big"
)

func swapEndian(in uint64) uint64 {
	out := uint64(0)
	for i := uint(0); i < 8; i++ {
		out |= ((in >> (8 * i)) & 0xff) << (56 - 8*i)
	}
	return out
}

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [2]uint64

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y baseFieldElem // value is x+yi.
}

func newGFp2() *gfP2 {
	return &gfP2{}
}

func (e *gfP2) String() string {
	return fmt.Sprintf("[%v, %v]", e.x, e.y)
}

func (e *gfP2) Int() *big.Int {
	return new(big.Int).SetBits([]big.Word{
		big.Word(swapEndian(e.y[1])),
		big.Word(swapEndian(e.y[0])),
		big.Word(swapEndian(e.x[1])),
		big.Word(swapEndian(e.x[0])),
	})
}

func (e *gfP2) Set(a *gfP2) *gfP2 {
	e.x[0], e.x[1] = a.x[0], a.x[1]
	e.y[0], e.y[1] = a.y[0], a.y[1]
	return e
}

func (e *gfP2) SetInt(in *big.Int) *gfP2 {
	w := in.Bits()
	if len(w) != 4 {
		return nil
	}

	e.y[1] = swapEndian(uint64(w[0]))
	e.y[0] = swapEndian(uint64(w[1]))
	e.x[1] = swapEndian(uint64(w[2]))
	e.x[0] = swapEndian(uint64(w[3]))

	return e
}

func (e *gfP2) SetZero() *gfP2 {
	e.x[0], e.x[1] = 0, 0
	e.y[0], e.y[1] = 0, 0
	return e
}

func (e *gfP2) SetOne() *gfP2 {
	e.x[0], e.x[1] = 1, 0
	e.y[0], e.y[1] = 0, 0
	return e
}

func (e *gfP2) IsZero() bool {
	return e.x[0] == 0 && e.x[1] == 0 && e.y[0] == 0 && e.y[1] == 0
}

//go:noescape
func feAdd(c, a, b *gfP2)

//go:noescape
func feSub(c, a, b *gfP2)

//go:noescape
func feMul(c, a, b *gfP2)

//go:noescape
func feSquare(c, a *gfP2)

//go:noescape
func feInvert(c, a *gfP2)

// reduce will set e.x or e.y to zero if they're equal to p. This is the only
// case where they will not naturally be reduced to canonical form.
func (e *gfP2) reduce() {
	if e.x[0] == bMask && e.x[1] == aMask {
		e.x[0], e.x[1] = 0, 0
	}
	if e.y[0] == bMask && e.y[1] == aMask {
		e.y[0], e.y[1] = 0, 0
	}
}
