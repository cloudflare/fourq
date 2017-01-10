// Package fourq implements FourQ, a high-speed elliptic curve at the 128-bit
// security level.
//
// https://eprint.iacr.org/2015/565.pdf
package fourq

func multByCofactor(pt *point) {
	temp := (&point{}).Set(pt)
	feMul(&temp.t, &temp.t, d)

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

// IsOnCurve returns true if the given is a point on the curve (including the
// identity point and points in a non-prime order subgroup) and false otherwise.
func IsOnCurve(in [32]byte) bool {
	_, ok := newPoint().SetBytes(in)
	return ok
}

// ScalarMult returns the point multiplied by scalar k.
func ScalarMult(in [32]byte, k []byte, clearCofactor bool) ([32]byte, bool) {
	pt, ok := (&point{}).SetBytes(in)
	if !ok {
		return [32]byte{}, false
	}

	if clearCofactor {
		multByCofactor(pt)
		pt.MakeAffine()
	}
	feMul(&pt.t, &pt.t, d)

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

	out := sum.Bytes()
	return out, out != [32]byte{1}
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
