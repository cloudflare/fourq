// +build !amd64 noasm

package fourq

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func feMul(e, a, b *gfP2) {
	tx, t := newBaseFieldElem(), newBaseFieldElem()
	bfeMul(tx, &a.x, &b.x)
	bfeMul(t, &a.y, &b.y)
	bfeSub(tx, tx, t)

	ty := newBaseFieldElem()
	bfeMul(ty, &a.x, &b.y)
	bfeMul(t, &a.y, &b.x)
	bfeAdd(ty, ty, t)

	e.x.Set(tx)
	e.y.Set(ty)
}

func feSquare(c, a *gfP2) {
	// Complex squaring algorithm:
	// (x+yi)² = (x+y)(x-y) + 2*x*y*i
	t1, t2, tx := newBaseFieldElem(), newBaseFieldElem(), newBaseFieldElem()
	bfeSub(t1, &a.x, &a.y)
	bfeAdd(t2, &a.x, &a.y)
	bfeMul(tx, t1, t2)

	bfeMul(t1, &a.x, &a.y)
	bfeDbl(t1, t1)

	c.x.Set(tx)
	c.y.Set(t1)
}
