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
		x: newGFp2(nil).SetZero(),
		y: newGFp2(nil).SetOne(),
		t: newGFp2(nil).SetZero(),
		z: newGFp2(nil).SetOne(),
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
	c.t.Mul(c.x, c.y, nil)
	c.z.SetOne()
	return c
}

func (c *point) Int(pool *elemPool) (x, y *big.Int) {
	c.MakeAffine(pool)

	x = new(big.Int).SetBytes(c.x.Bytes())
	y = new(big.Int).SetBytes(c.y.Bytes())
	return
}

func (c *point) IsOnCurve(pool *elemPool) bool {
	z2 := newGFp2(pool).Square(c.z, pool)
	z4 := newGFp2(pool).Square(z2, pool)

	x2 := newGFp2(pool).Square(c.x, pool)
	y2 := newGFp2(pool).Square(c.y, pool)

	lhs := newGFp2(pool).Sub(y2, x2)
	lhs.Mul(lhs, z2, pool)

	rhs := newGFp2(pool).Square(c.t, pool)
	rhs.Mul(rhs, d, pool).Add(rhs, z4)

	ok := lhs.Sub(lhs, rhs).IsZero()

	z2.Put(pool)
	z4.Put(pool)
	x2.Put(pool)
	y2.Put(pool)
	lhs.Put(pool)
	rhs.Put(pool)

	return ok
}

func (c *point) Add(a, b *point, pool *elemPool) *point {
	A := newGFp2(pool).Sub(a.y, a.x)
	tmp := newGFp2(pool).Add(b.y, b.x)
	A.Mul(A, tmp, pool)

	B := newGFp2(pool).Add(a.y, a.x)
	tmp.Sub(b.y, b.x)
	B.Mul(B, tmp, pool)

	C := newGFp2(pool).Mul(a.z, b.t, pool)
	C.Dbl(C)

	D := newGFp2(pool).Mul(a.t, b.z, pool)
	D.Dbl(D)

	E := newGFp2(pool).Add(D, C)
	F := newGFp2(pool).Sub(B, A)
	G := newGFp2(pool).Add(B, A)
	H := newGFp2(pool).Sub(D, C)

	c.x.Mul(E, F, pool)
	c.y.Mul(G, H, pool)
	c.t.Mul(E, H, pool)
	c.z.Mul(F, G, pool)

	tmp.Put(pool)
	A.Put(pool)
	B.Put(pool)
	C.Put(pool)
	D.Put(pool)
	E.Put(pool)
	F.Put(pool)
	G.Put(pool)
	H.Put(pool)

	return c
}

func (c *point) Dbl(a *point, pool *elemPool) *point {
	A := newGFp2(pool).Square(a.x, pool)
	B := newGFp2(pool).Square(a.y, pool)
	C := newGFp2(pool).Square(a.z, pool)
	C.Dbl(C)

	// D = negative A

	E := newGFp2(pool).Add(a.x, a.y)
	E.Square(E, pool).Sub(E, A).Sub(E, B)
	G := newGFp2(pool).Sub(B, A)
	F := newGFp2(pool).Sub(G, C)
	H := newGFp2(pool).Add(A, B)
	H.Neg(H)

	c.x.Mul(E, F, pool)
	c.y.Mul(G, H, pool)
	c.t.Mul(E, H, pool)
	c.z.Mul(F, G, pool)

	A.Put(pool)
	B.Put(pool)
	C.Put(pool)
	E.Put(pool)
	G.Put(pool)
	F.Put(pool)
	H.Put(pool)

	return c
}

func (c *point) MakeAffine(pool *elemPool) {
	zInv := newGFp2(pool).Invert(c.z, pool)

	c.x.Mul(c.x, zInv, pool)
	c.y.Mul(c.y, zInv, pool)
	c.t.Mul(c.t, zInv, pool)
	c.t.Mul(c.x, c.y, pool)
	c.z.SetOne()

	zInv.Put(pool)
}
