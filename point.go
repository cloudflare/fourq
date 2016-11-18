package fourq

import (
	"fmt"
	"math/big"
)

type point struct {
	x, y, t, z gfP2
}

func newPoint() *point {
	pt := &point{}
	pt.y.SetOne()
	pt.z.SetOne()
	return pt
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
	if c == nil {
		return false
	}

	z2, z4 := newGFp2(), newGFp2()
	feSquare(z2, &c.z)
	feSquare(z4, z2)

	x2, y2 := newGFp2(), newGFp2()
	feSquare(x2, &c.x)
	feSquare(y2, &c.y)

	lhs := newGFp2()
	feSub(lhs, y2, x2)
	feMul(lhs, lhs, z2)

	rhs := newGFp2()
	feSquare(rhs, &c.t)
	feMul(rhs, rhs, d)
	feAdd(rhs, rhs, z4)

	feSub(lhs, lhs, rhs)
	lhs.reduce()
	return lhs.IsZero()
}

//go:noescape
func pMixedAdd(a, b *point)

//go:noescape
func pDbl(a *point)

func (c *point) MakeAffine() {
	// zInv := newGFp2().Invert(c.z)
	feInvert(&c.z, &c.z)

	feMul(&c.x, &c.x, &c.z)
	feMul(&c.y, &c.y, &c.z)
	// feMul(c.t, c.t, zInv)
	// feMul(c.z, c.x, c.y)
	// c.z.SetOne()

	c.x.reduce()
	c.y.reduce()
	// c.t.reduce()
}
