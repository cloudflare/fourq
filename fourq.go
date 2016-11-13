// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
package fourq

import (
	"math/big"
)

func IsOnCurve(xI, yI *big.Int) bool {
	if xI == nil || yI == nil {
		return false
	}

	pt := newPoint()
	pt.SetInt(xI, yI)
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

func ScalarBaseMult(k []byte) (x, y *big.Int) {
	if len(k) == 0 {
		return nil, nil
	}

	sum := &point{}
	sum.Set(generatorBase[k[0]])

	for _, byte := range k[1:] { // TODO(brendan): Mult by cofactor.
		for bit := 0; bit < 8; bit++ {
			pDbl(sum)
		}
		pMixedAdd(sum, generatorBase[byte])
	}

	return sum.Int()
}
