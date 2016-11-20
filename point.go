package fourq

import (
	"fmt"
	"math/big"
)

type point struct {
	x, y, t, z fieldElem
}

func newPoint() *point {
	pt := &point{}
	pt.y.SetOne()
	pt.z.SetOne()
	return pt
}

func decompressPoint(y *big.Int) (*point, bool) {
	pt := &point{}
	pt.y.SetInt(y)
	pt.z.SetOne()

	// Separate p.y from the sign of x.
	var s uint64
	s, pt.y.y[1] = uint64(pt.y.y[1])>>63, pt.y.y[1]&aMask

	if pt.y.x[1]>>63 == 1 {
		return nil, false
	}

	u, v := newFieldElem(), newFieldElem()
	feSquare(u, &pt.y)
	feSub(u, u, one)

	feSquare(v, &pt.y)
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

	x0 := newBaseFieldElem()
	bfeHalf(x0, b)

	x1 := newBaseFieldElem()
	bfeMul(x1, a, t2)
	bfeMul(x1, x1, t1)

	bfeSquare(temp, b)
	bfeMul(temp, temp, t2)
	if *temp != *t {
		x0, x1 = x1, x0
	}

	pt.x = fieldElem{x: *x0, y: *x1}
	if pt.x.sign() != s {
		pt.x.Neg(&pt.x)
	}

	if !pt.IsOnCurve() {
		pt.x.y.Neg(&pt.x.y)
	}
	if !pt.IsOnCurve() {
		return nil, false
	}
	return pt, true
}

func (c *point) String() string {
	return fmt.Sprintf("point(\n\tx: %v,\n\ty: %v,\n\tt: %v,\n\tz: %v\n)", &c.x, &c.y, &c.t, &c.z)
}

func (c *point) Set(a *point) *point {
	c.x.Set(&a.x)
	c.y.Set(&a.y)
	c.t.Set(&a.t)
	c.z.Set(&a.z)
	return c
}

func (c *point) SetInt(x, y *big.Int) *point {
	c.x.SetInt(x)
	c.y.SetInt(y)
	feMul(&c.t, &c.x, &c.y)
	c.z.SetOne()
	return c
}

func (c *point) Int() (x, y *big.Int) {
	c.MakeAffine()
	return c.x.Int(), c.y.Int()
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

func pMixedAdd(a, b *point) {
	A := newFieldElem()
	feMul(A, &a.x, &b.x)

	B := newFieldElem()
	feMul(B, &a.y, &b.y)

	C := newFieldElem()
	feMul(C, &a.t, &b.t)
	feMul(C, C, d)

	D := newFieldElem()
	feMul(D, &a.z, &b.z)

	E, t := newFieldElem(), newFieldElem()
	feAdd(E, &a.x, &a.y)
	feAdd(t, &b.x, &b.y)
	feMul(E, E, t)
	feSub(E, E, A)
	feSub(E, E, B)

	F := newFieldElem()
	feSub(F, D, C)

	G := newFieldElem()
	feAdd(G, D, C)

	H := newFieldElem()
	feAdd(H, A, B)

	feMul(&a.x, E, F)
	feMul(&a.y, G, H)
	feMul(&a.t, E, H)
	feMul(&a.z, F, G)
}

//go:noescape
func pDbl(a *point)

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
