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

// fieldElem implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type fieldElem struct {
	x, y baseFieldElem // value is x+yi.
}

func newFieldElem() *fieldElem {
	return &fieldElem{}
}

func (e *fieldElem) String() string {
	return fmt.Sprintf("[%v, %v]", e.x, e.y)
}

func (e *fieldElem) Int() *big.Int {
	return new(big.Int).SetBits([]big.Word{
		big.Word(swapEndian(e.y[1])),
		big.Word(swapEndian(e.y[0])),
		big.Word(swapEndian(e.x[1])),
		big.Word(swapEndian(e.x[0])),
	})
}

func (e *fieldElem) Set(a *fieldElem) *fieldElem {
	e.x[0], e.x[1] = a.x[0], a.x[1]
	e.y[0], e.y[1] = a.y[0], a.y[1]
	return e
}

func (e *fieldElem) SetInt(in *big.Int) *fieldElem {
	w, temp := [4]uint64{}, in.Bits()
	for i := 0; i < len(temp) && i < 4; i++ {
		w[i] = uint64(temp[i])
	}

	e.y[1] = swapEndian(w[0])
	e.y[0] = swapEndian(w[1])
	e.x[1] = swapEndian(w[2])
	e.x[0] = swapEndian(w[3])

	return e
}

func (e *fieldElem) SetZero() *fieldElem {
	e.x[0], e.x[1] = 0, 0
	e.y[0], e.y[1] = 0, 0
	return e
}

func (e *fieldElem) SetOne() *fieldElem {
	e.x[0], e.x[1] = 1, 0
	e.y[0], e.y[1] = 0, 0
	return e
}

func (e *fieldElem) IsZero() bool {
	return e.x[0] == 0 && e.x[1] == 0 && e.y[0] == 0 && e.y[1] == 0
}

//go:noescape
func feAdd(c, a, b *fieldElem)

//go:noescape
func feSub(c, a, b *fieldElem)

//go:noescape
func feMul(c, a, b *fieldElem)

//go:noescape
func feSquare(c, a *fieldElem)

//go:noescape
func feInvert(c, a *fieldElem)

// reduce will set e.x or e.y to zero if they're equal to p. This is the only
// case where they will not naturally be reduced to canonical form.
func (e *fieldElem) reduce() {
	if e.x[0] == bMask && e.x[1] == aMask {
		e.x[0], e.x[1] = 0, 0
	}
	if e.y[0] == bMask && e.y[1] == aMask {
		e.y[0], e.y[1] = 0, 0
	}
}
