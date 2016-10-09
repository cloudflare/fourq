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
	c.t.Mul(c.x, c.y)
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
	z2 := newGFp2().Square(c.z)
	z4 := newGFp2().Square(z2)

	x2 := newGFp2().Square(c.x)
	y2 := newGFp2().Square(c.y)

	lhs := newGFp2().Sub(y2, x2)
	lhs.Mul(lhs, z2)

	rhs := newGFp2().Square(c.t)
	rhs.Mul(rhs, d).Add(rhs, z4)

	ok := lhs.Sub(lhs, rhs).IsZero()

	return ok
}

func (c *point) Add(a, b *point) *point {
	A := newGFp2().Sub(a.y, a.x)
	tmp := newGFp2().Add(b.y, b.x)
	A.Mul(A, tmp)

	B := newGFp2().Add(a.y, a.x)
	tmp.Sub(b.y, b.x)
	B.Mul(B, tmp)

	C := newGFp2().Mul(a.z, b.t)
	C.Dbl(C)

	D := newGFp2().Mul(a.t, b.z)
	D.Dbl(D)

	E := newGFp2().Add(D, C)
	F := newGFp2().Sub(B, A)
	G := newGFp2().Add(B, A)
	H := newGFp2().Sub(D, C)

	c.x.Mul(E, F)
	c.y.Mul(G, H)
	c.t.Mul(E, H)
	c.z.Mul(F, G)

	return c
}

func (c *point) Dbl(a *point) *point {
	A := newGFp2().Square(a.x)
	B := newGFp2().Square(a.y)
	C := newGFp2().Square(a.z)
	C.Dbl(C)

	// D = negative A

	E := newGFp2().Add(a.x, a.y)
	E.Square(E).Sub(E, A).Sub(E, B)
	G := newGFp2().Sub(B, A)
	F := newGFp2().Sub(G, C)
	H := newGFp2().Add(A, B)
	H.Neg(H)

	c.x.Mul(E, F)
	c.y.Mul(G, H)
	c.t.Mul(E, H)
	c.z.Mul(F, G)

	return c
}

func (c *point) MakeAffine() {
	zInv := newGFp2().Invert(c.z)

	c.x.Mul(c.x, zInv)
	c.y.Mul(c.y, zInv)
	c.t.Mul(c.t, zInv)
	c.t.Mul(c.x, c.y)
	c.z.SetOne()
}
