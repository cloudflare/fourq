// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
package fourq

import (
	"math/big"
)

func clearCofactor(pt *point) {
	temp := (&point{}).Set(pt)

	pDbl(pt)
	pMixedAdd(pt, temp)
	pDbl(pt)
	pDbl(pt)
	pDbl(pt)
	pDbl(pt)
	pMixedAdd(pt, temp)
	pDbl(pt)
	pDbl(pt)
	pDbl(pt)
}

// IsOnCurve returns true if (xI, yI) is a point on the curve (including the
// identity point and points in a non-prime order subgroup) and false otherwise.
func IsOnCurve(xI, yI *big.Int) bool {
	if xI == nil || yI == nil {
		return false
	}

	_, ok := newPoint().SetInt(xI, yI)
	return ok
}

func ScalarMult(xI, yI *big.Int, k []byte) (*big.Int, *big.Int) {
	if xI == nil || yI == nil {
		return nil, nil
	}

	pt, ok := (&point{}).SetInt(xI, yI)
	if !ok {
		return nil, nil
	}
	feMul(&pt.t, &pt.t, d)
	// TODO(brendan): Mult by cofactor

	sum := newPoint()

	for _, byte := range k {
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
