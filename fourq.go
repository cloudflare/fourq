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

	return pt.IsOnCurve()
}

func ScalarMult(xI, yI *big.Int, k []byte) (*big.Int, *big.Int) {
	pt := newPoint()
	pt.SetBytes(unpack(xI), unpack(yI))

	// TODO(brendan): Check if point is on curve?
	// TODO(brendan): Mult by cofactor.

	sum := newPoint()
	tmp := newPoint()

	for pos := 0; pos < 256; pos++ { // TODO(brendan): Allow exp larger than order?
		b := k[pos/8] >> uint(7-(pos%8))
		b &= 1

		tmp.Dbl(sum)
		if b == 1 {
			sum.Add(tmp, pt)
		} else {
			sum.Set(tmp)
		}
	}

	return sum.Int()
}

func ScalarBaseMult(k []byte) (x, y *big.Int) { return ScalarMult(Gx, Gy, k) }
