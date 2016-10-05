package fourq

import (
	"math/big"
)

func unpack(in *big.Int) []byte {
	out := make([]byte, 32)
	inBytes := in.Bytes()
	if len(inBytes) > 32 {
		return out
	}

	copy(out[32-len(inBytes):], inBytes)
	return out
}

func IsOnCurve(xI, yI *big.Int) bool {
	pt := newPoint()
	pt.SetBytes(unpack(xI), unpack(yI))

	return pt.IsOnCurve(nil)
}

func ScalarMult(xI, yI *big.Int, k []byte) (*big.Int, *big.Int) {
	pt := newPoint()
	pt.SetBytes(unpack(xI), unpack(yI))

	// TODO(brendan): Check if point is on curve?
	// TODO(brendan): Mult by cofactor.

	pool := new(elemPool)
	sum := newPoint()
	tmp := newPoint()

	for pos := 0; pos < 256; pos++ { // TODO(brendan): Allow exp larger than order?
		b := k[pos/8] >> uint(7-(pos%8))
		b &= 1

		tmp.Dbl(sum, pool)
		if b == 1 {
			sum.Add(tmp, pt, pool)
		} else {
			sum.Set(tmp)
		}
	}

	return sum.Int(pool)
}

func ScalarBaseMult(k []byte) (x, y *big.Int) { return ScalarMult(Gx, Gy, k) }

// elemPool implements a tiny cache of *baseFieldElem objects that's used to
// reduce the number of allocations made during processing.
type elemPool struct {
	elems []*baseFieldElem
	// count int
}

func (pool *elemPool) Get() *baseFieldElem {
	if pool == nil {
		return &baseFieldElem{}
	}
	// pool.count += 1

	l := len(pool.elems)
	if l == 0 {
		return &baseFieldElem{}
	}

	elem := pool.elems[l-1]
	pool.elems = pool.elems[:l-1]
	return elem
}

func (pool *elemPool) Put(elem *baseFieldElem) {
	if pool == nil {
		return
	}
	// pool.count -= 1
	pool.elems = append(pool.elems, elem)
}
