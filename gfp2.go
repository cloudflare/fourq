package fourq

import (
	"fmt"
)

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y baseFieldElem // value is x+yi.
}

func newGFp2() *gfP2 {
	return &gfP2{*newBaseFieldElem(), *newBaseFieldElem()}
}

func (e *gfP2) String() string {
	return fmt.Sprintf("[%v, %v]", e.x, e.y)
}

func (e *gfP2) Bytes() []byte {
	ret := make([]byte, 32)
	copy(ret[:16], e.x.Bytes())
	copy(ret[16:], e.y.Bytes())
	return ret
}

func (e *gfP2) Set(a *gfP2) *gfP2 {
	e.x.Set(&a.x)
	e.y.Set(&a.y)
	return e
}

func (e *gfP2) SetBytes(in []byte) *gfP2 {
	if len(in) != 32 {
		return nil
	}

	e.x.SetBytes(in[:16])
	e.y.SetBytes(in[16:])
	return e
}

func (e *gfP2) SetZero() *gfP2 {
	e.x.SetZero()
	e.y.SetZero()
	return e
}

func (e *gfP2) SetOne() *gfP2 {
	e.x.SetOne()
	e.y.SetZero()
	return e
}

func (e *gfP2) IsZero() bool {
	return e.x.IsZero() && e.y.IsZero()
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

func (e *gfP2) reduce() {
	e.x.reduce()
	e.y.reduce()
}
