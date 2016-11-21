package fourq

import (
	"fmt"
	"math/big"
)

type point struct {
	x, y, z, t fieldElem
}

func newPoint() *point {
	pt := &point{}
	pt.y.SetOne()
	pt.z.SetOne()
	return pt
}

func (c *point) String() string {
	return fmt.Sprintf("point(\n\tx: %v,\n\ty: %v,\n\tz: %v,\n\tt: %v\n)", &c.x, &c.y, &c.z, &c.t)
}

func (c *point) Set(a *point) *point {
	c.x.Set(&a.x)
	c.y.Set(&a.y)
	c.t.Set(&a.t)
	c.z.Set(&a.z)
	return c
}

func (c *point) SetInt(x, y *big.Int) (*point, bool) {
	c.y.x.SetInt(x)
	c.y.y.SetInt(y)
	c.z.SetOne()

	// Separate p.y from the sign of x.
	var s uint64
	s, c.y.y[1] = uint64(c.y.y[1])>>63, c.y.y[1]&aMask

	if c.y.x[1]>>63 == 1 {
		return nil, false
	}

	// Recover x coordinate from y, up to a multiple of plus/minus one.
	u, v := newFieldElem(), newFieldElem()
	feSquare(u, &c.y)
	feSub(u, u, one)

	feSquare(v, &c.y)
	feMul(v, v, d)
	feAdd(v, v, one)

	t0, temp := newBaseFieldElem(), newBaseFieldElem()
	bfeMul(t0, &u.x, &v.x)
	bfeMul(temp, &u.y, &v.y)
	bfeAdd(t0, t0, temp)

	t1 := newBaseFieldElem()
	bfeMul(t1, &u.y, &v.x)
	bfeMul(temp, &u.x, &v.y)
	bfeSub(t1, temp, t1)

	t2 := newBaseFieldElem()
	bfeSquare(t2, &v.x)
	bfeSquare(temp, &v.y)
	bfeAdd(t2, t2, temp)

	t3 := newBaseFieldElem()
	bfeSquare(t3, t0)
	bfeSquare(temp, t1)
	bfeAdd(t3, t3, temp)
	for i := 0; i < 125; i++ {
		bfeSquare(t3, t3)
	}

	t := newBaseFieldElem()
	bfeAdd(t, t0, t3)
	t.reduce()
	if t.IsZero() {
		bfeSub(t, t0, t3)
	}
	bfeDbl(t, t)

	a := newBaseFieldElem()
	bfeSquare(a, t2)
	bfeMul(a, a, t2)
	bfeMul(a, a, t)
	a.chain1251(a)

	b := newBaseFieldElem()
	bfeMul(b, a, t2)
	bfeMul(b, b, t)

	bfeHalf(&c.x.x, b)
	bfeMul(&c.x.y, a, t2)
	bfeMul(&c.x.y, &c.x.y, t1)

	// Recover x-coordinate exactly.
	bfeSquare(temp, b)
	bfeMul(temp, temp, t2)
	if *temp != *t {
		c.x.x, c.x.y = c.x.y, c.x.x
	}
	if c.x.sign() != s {
		c.x.Neg(&c.x)
	}
	if !c.IsOnCurve() {
		c.x.y.Neg(&c.x.y)
	}

	// Finally, verify point is valid and return.
	if !c.IsOnCurve() {
		return nil, false
	}

	feMul(&c.t, &c.x, &c.y)
	return c, true
}

func (c *point) Int() (x, y *big.Int) {
	c.MakeAffine()
	c.y.y[1] += c.x.sign() << 63
	return c.y.x.Int(), c.y.y.Int()
}

func (c *point) IsOnCurve() bool {
	x2, y2 := newFieldElem(), newFieldElem()
	feSquare(x2, &c.x)
	feSquare(y2, &c.y)

	lhs := newFieldElem()
	feSub(lhs, y2, x2)

	rhs := newFieldElem()
	feMul(rhs, &c.x, &c.y)
	feSquare(rhs, rhs)
	feMul(rhs, rhs, d)
	feAdd(rhs, rhs, one)

	feSub(lhs, lhs, rhs)
	lhs.reduce()
	return lhs.IsZero()
}

func (c *point) MakeAffine() {
	// zInv := newFieldElem().Invert(c.z)
	c.z.Invert(&c.z)

	feMul(&c.x, &c.x, &c.z)
	feMul(&c.y, &c.y, &c.z)
	// feMul(c.t, c.t, zInv)
	// feMul(c.z, c.x, c.y)
	// c.z.SetOne()

	c.x.reduce()
	c.y.reduce()
	// c.t.reduce()
}

//go:noescape
func pDbl(a *point)

func pMixedAdd(a, b *point) {
	A := newFieldElem()
	feMul(A, &a.x, &b.x)

	B := newFieldElem()
	feMul(B, &a.y, &b.y)

	C := newFieldElem()
	feMul(C, &a.t, &b.t)
	feMul(C, C, d)

	// D = Z1

	E, temp := newFieldElem(), newFieldElem()
	feAdd(E, &a.x, &a.y)
	feAdd(temp, &b.x, &b.y)
	feMul(E, E, temp)
	feSub(E, E, A)
	feSub(E, E, B)

	F := newFieldElem()
	feSub(F, &a.z, C)

	G := newFieldElem()
	feAdd(G, &a.z, C)

	H := newFieldElem()
	feAdd(H, A, B)

	feMul(&a.x, E, F)
	feMul(&a.y, G, H)
	feMul(&a.z, F, G)
	feMul(&a.t, E, H)
}
