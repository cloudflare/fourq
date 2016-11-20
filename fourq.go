// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
package fourq

import (
	"math/big"
)

// IsOnCurve returns true if (xI, yI) is a point on the curve (including the
// identity point and points in a non-prime order subgroup) and false otherwise.
func IsOnCurve(xI, yI *big.Int) bool {
	if xI == nil || yI == nil {
		return false
	}

	pt := newPoint().SetInt(xI, yI)
	return pt.IsOnCurve()
}

func ScalarMult(xI, yI *big.Int, k []byte) (*big.Int, *big.Int) {
	if xI == nil || yI == nil {
		return nil, nil
	}

	pt := &point{}
	// TODO(brendan): Check if point is on curve?
	pt.SetInt(xI, yI) // TODO(brendan): Point decompression.

	sum := newPoint()

	for _, byte := range k { // TODO(brendan): Mult by cofactor.
		for bit := 0; bit < 8; bit++ {
			pDbl(sum)
			if byte&0x80 == 0x080 {
				pMixedAdd(sum, pt)
			}
			byte <<= 1
		}
	}

	return sum.Int()
}

func ScalarBaseMult(k []byte) (x, y *big.Int) { return ScalarMult(Gx, Gy, k) }
