package fourq

import (
	"fmt"
	"math/big"
)

type point struct {
	x, y, t, z *gfP2
}

func newPoint() *point {
	return &point{
		x: newGFp2().SetZero(),
		y: newGFp2().SetOne(),
		t: newGFp2().SetZero(),
		z: newGFp2().SetOne(),
	}
}

func (c *point) String() string {
	return fmt.Sprintf("point(\n\tx: %v,\n\ty: %v,\n\tt: %v,\n\tz: %v\n)", c.x, c.y, c.t, c.z)
}

func (c *point) Set(a *point) *point {
	c.x.Set(a.x)
	c.y.Set(a.y)
	c.t.Set(a.t)
	c.z.Set(a.z)
	return c
}

func (c *point) SetBytes(x, y []byte) *point {
	c.x.SetBytes(x)
	c.y.SetBytes(y)
	feMul(c.t, c.x, c.y)
	c.z.SetOne()
	return c
}

func (c *point) Int() (x, y *big.Int) {
	c.MakeAffine()

	x = new(big.Int).SetBytes(c.x.Bytes())
	y = new(big.Int).SetBytes(c.y.Bytes())
	return
}

func (c *point) IsOnCurve() bool {
	z2, z4 := newGFp2(), newGFp2()
	feSquare(z2, c.z)
	feSquare(z4, z2)

	x2, y2 := newGFp2(), newGFp2()
	feSquare(x2, c.x)
	feSquare(y2, c.y)

	lhs := newGFp2().Sub(y2, x2)
	feMul(lhs, lhs, z2)

	rhs := newGFp2()
	feSquare(rhs, c.t)
	feMul(rhs, rhs, d)
	rhs.Add(rhs, z4)

	lhs.Sub(lhs, rhs).reduce()
	return lhs.IsZero()
}

func (c *point) Add(a, b *point) *point {
	A := newGFp2().Sub(a.y, a.x)
	tmp := newGFp2().Add(b.y, b.x)
	feMul(A, A, tmp)

	B := newGFp2().Add(a.y, a.x)
	tmp.Sub(b.y, b.x)
	feMul(B, B, tmp)

	C := newGFp2()
	feMul(C, a.z, b.t)
	C.Dbl(C)

	D := newGFp2()
	feMul(D, a.t, b.z)
	D.Dbl(D)

	E := newGFp2().Add(D, C)
	F := newGFp2().Sub(B, A)
	G := newGFp2().Add(B, A)
	H := newGFp2().Sub(D, C)

	feMul(c.x, E, F)
	feMul(c.y, G, H)
	feMul(c.t, E, H)
	feMul(c.z, F, G)

	return c
}

func (c *point) Dbl(a *point) *point {
	A, B, C := newGFp2(), newGFp2(), newGFp2()
	feSquare(A, a.x)
	feSquare(B, a.y)
	feSquare(C, a.z)
	C.Dbl(C)

	// D = negative A

	E := newGFp2().Add(a.x, a.y)
	feSquare(E, E)
	E.Sub(E, A).Sub(E, B)
	G := newGFp2().Sub(B, A)
	F := newGFp2().Sub(G, C)
	H := newGFp2().Add(A, B)
	H.Neg(H)

	feMul(c.x, E, F)
	feMul(c.y, G, H)
	feMul(c.t, E, H)
	feMul(c.z, F, G)

	return c
}

func (c *point) MakeAffine() {
	// zInv := newGFp2().Invert(c.z)
	c.z.Invert(c.z)

	feMul(c.x, c.x, c.z)
	feMul(c.y, c.y, c.z)
	// feMul(c.t, c.t, zInv)
	// feMul(c.z, c.x, c.y)
	// c.z.SetOne()

	c.x.reduce()
	c.y.reduce()
	// c.t.reduce()
}
