// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
package fourq

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

// IsOnCurve returns true if pt is a point on the curve (including the identity
// point and points in a non-prime order subgroup) and false otherwise.
func IsOnCurve(in [32]byte) bool {
	_, ok := newPoint().SetBytes(in)
	return ok
}

// ScalarMult returns the point multiplied by scalar k.
func ScalarMult(in [32]byte, k []byte) [32]byte {
	pt, ok := (&point{}).SetBytes(in)
	if !ok {
		return [32]byte{}
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

	return sum.Bytes()
}

// ScalarBaseMult returns the generator multiplied by scalar k. k's slice should
// be 32 bytes long or shorter.
func ScalarBaseMult(k []byte) [32]byte {
	if len(k) > 32 {
		return [32]byte{}
	}
	K := make([]byte, 32)
	copy(K[32-len(k):], k)

	sum := newPoint()

	for i := 0; i < 4; i++ {
		for bit := uint(0); bit < 8; bit++ {
			var idx byte
			for block := 0; block < 8; block++ {
				idx = 2*idx + ((K[4*block+i] >> (7 - bit)) & 1)
			}

			pDbl(sum)
			if idx != 0 {
				pMixedAdd(sum, generatorBase[idx])
			}
		}
	}

	return sum.Bytes()
}
