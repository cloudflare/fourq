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

// ScalarMult returns the point (xI, yI) multiplied by scalar k.
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

// ScalarBaseMult returns the generator multiplied by scalar k. k's slice should
// be 32 bytes long or shorter.
func ScalarBaseMult(k []byte) (x, y *big.Int) {
	if len(k) > 32 {
		return nil, nil
	}
	K := make([]byte, 32)
	copy(k, K[32-len(k):])

	sum := newPoint()

	for i := 0; i < 16; i++ {
		for bit := uint(0); bit < 8; bit++ {
			idx := (K[31-i] >> (7 - bit)) & 1
			idx = 2*idx + (K[15-i]>>(7-bit))&1

			pDbl(sum)
			if idx != 0 {
				pMixedAdd(sum, generatorBase[idx])
			}
		}
	}

	return sum.Int()
}
