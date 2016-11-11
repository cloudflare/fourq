// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
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

	for _, byte := range k { // TODO(brendan): Allow exp larger than Order?
		for bitNum := 0; bitNum < 8; bitNum++ {
			sum.Dbl(sum)
			if byte&0x80 == 0x080 {
				sum.MixedAdd(sum, pt)
			}
			byte <<= 1
		}
	}

	return sum.Int()
}

func ScalarBaseMult(k []byte) (x, y *big.Int) { return ScalarMult(Gx, Gy, k) }
