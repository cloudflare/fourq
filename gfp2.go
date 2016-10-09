package fourq

import (
	"fmt"
	"math/big"
)

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y *baseFieldElem // value is x+yi.
}

func newGFp2() *gfP2 {
	return &gfP2{newBaseFieldElem(), newBaseFieldElem()}
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
	e.x.Set(a.x)
	e.y.Set(a.y)
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

func (e *gfP2) Neg(a *gfP2) *gfP2 {
	e.x.Neg(a.x)
	e.y.Neg(a.y)
	return e
}

func (e *gfP2) Dbl(a *gfP2) *gfP2 {
	bfeDbl(e.x, a.x)
	bfeDbl(e.y, a.y)
	return e
}

func (e *gfP2) Add(a, b *gfP2) *gfP2 {
	bfeAdd(e.x, a.x, b.x)
	bfeAdd(e.y, a.y, b.y)
	return e
}

func (e *gfP2) Sub(a, b *gfP2) *gfP2 {
	bfeSub(e.x, a.x, b.x)
	bfeSub(e.y, a.y, b.y)
	return e
}

func (c *gfP2) Exp(a *gfP2, power *big.Int) *gfP2 {
	sum := newGFp2()
	sum.SetOne()
	t := newGFp2()

	for i := power.BitLen() - 1; i >= 0; i-- {
		t.Square(sum)
		if power.Bit(i) != 0 {
			sum.Mul(t, a)
		} else {
			sum.Set(t)
		}
	}

	c.Set(sum)
	return c
}

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func (e *gfP2) Mul(a, b *gfP2) *gfP2 {
	tx, t := newBaseFieldElem(), newBaseFieldElem()
	bfeMul(tx, a.x, b.x)
	bfeMul(t, a.y, b.y)
	bfeSub(tx, tx, t)

	ty := newBaseFieldElem()
	bfeMul(ty, a.x, b.y)
	bfeMul(t, a.y, b.x)
	bfeAdd(ty, ty, t)

	e.x.Set(tx)
	e.y.Set(ty)
	return e
}

func (e *gfP2) Square(a *gfP2) *gfP2 {
	// Complex squaring algorithm:
	// (x+yi)² = (x+y)(x-y) + 2*x*y*i
	t1, t2, tx := newBaseFieldElem(), newBaseFieldElem(), newBaseFieldElem()
	bfeSub(t1, a.x, a.y)
	bfeAdd(t2, a.x, a.y)
	bfeMul(tx, t1, t2)

	bfeMul(t1, a.x, a.y)
	bfeDbl(t1, t1)

	e.x.Set(tx)
	e.y.Set(t1)
	return e
}

func (e *gfP2) Invert(a *gfP2) *gfP2 {
	// See "Implementing cryptographic pairings", M. Scott, section 3.2.
	// ftp://136.206.11.249/pub/crypto/pairings.pdf
	t, t2 := newBaseFieldElem(), newBaseFieldElem()
	bfeSquare(t, a.x)
	bfeSquare(t2, a.y)
	bfeAdd(t, t, t2)

	inv := newBaseFieldElem().Invert(t)

	e.y.Neg(a.y)
	bfeMul(e.y, e.y, inv)
	bfeMul(e.x, a.x, inv)

	return e
}
