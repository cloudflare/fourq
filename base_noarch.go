package fourq

// The approach to modular arithmetic in this file is as follows: We represent a
// number mod p=2^127-1 by spreading its bits across 5 uint64s, where only the
// bottom few bits of each uint64 are occupied in a fully reduced element.
// Specifically, the bottom 26 bits of the first and third uint64, and the
// bottom 25 bits of every other. This was chosen because it makes element
// multiplication and squaring easy while avoiding overflow.
//
// Functions on a number which can be implemented substantially faster in
// assembly are addressed as bfeXXX(c, a, b *bfe) while functions which cannot
// are addressed as methods: (c *bfe).XXX(a, b *bfe) (c *bfe).

import (
	"fmt"
	"math/big"
)

var (
	aMask uint64 = 0x01ffffff
	bMask uint64 = 0x03ffffff

	oneBFE = &baseFieldElem{1, 0, 0, 0, 0}
)

func getBit(in []byte, k int) uint64 {
	return uint64(in[k/8]>>uint(k%8)) & 1
}

func aSplit(in uint64) (uint64, uint64) { return in & aMask, in >> 25 }
func bSplit(in uint64) (uint64, uint64) { return in & bMask, in >> 26 }

func aaSplit(in uint64) (uint64, uint64, uint64) {
	return in & aMask, (in >> 25) & aMask, in >> 50
}

func abSplit(in uint64) (uint64, uint64, uint64) {
	return in & aMask, (in >> 25) & bMask, in >> 51
}

func baSplit(in uint64) (uint64, uint64, uint64) {
	return in & bMask, (in >> 26) & aMask, in >> 51
}

func aNeg(in uint64) uint64 { return (^in) & aMask }
func bNeg(in uint64) uint64 { return (^in) & bMask }

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [5]uint64

func newBaseFieldElem() *baseFieldElem {
	return &baseFieldElem{}
}

// numToBFE takes a big.Int as input and returns its representation as a
// baseFieldElement. This function is used exclusively in tests.
func numToBFE(in *big.Int) *baseFieldElem {
	out := newBaseFieldElem()

	for i := 0; i < 26; i++ {
		out[0] += uint64(in.Bit(i)) << uint(i)
	}
	for i := 0; i < 25; i++ {
		out[1] += uint64(in.Bit(26+i)) << uint(i)
	}
	for i := 0; i < 26; i++ {
		out[2] += uint64(in.Bit(26+25+i)) << uint(i)
	}
	for i := 0; i < 25; i++ {
		out[3] += uint64(in.Bit(26+25+26+i)) << uint(i)
	}
	for i := 0; i < 25; i++ {
		out[4] += uint64(in.Bit(26+25+26+25+i)) << uint(i)
	}

	return out
}

func (c *baseFieldElem) String() string {
	return fmt.Sprintf("%x %x %x %x %x", c[0], c[1], c[2], c[3], c[4])
}

// Bytes returns the compressed, little-endian representation of a number.
// It is compatible with c.SetBytes().
func (c *baseFieldElem) Bytes() []byte {
	out := make([]byte, 16)

	for i := 0; i < 2; i++ {
		for j := uint(0); j < 26; j++ {
			k, b := (26+25)*i+int(j), byte(c[2*i+0]>>j)&1
			out[k/8] += b << uint(k%8)
		}
		for j := uint(0); j < 25; j++ {
			k, b := (26+25)*i+int(j)+26, byte(c[2*i+1]>>j)&1
			out[k/8] += b << uint(k%8)
		}
	}
	for j := uint(0); j < 25; j++ {
		k, b := int(j)+102, byte(c[4]>>j)&1
		out[k/8] += b << uint(k%8)
	}

	return out
}

// Set sets c to be a (with duplication) and returns c.
func (c *baseFieldElem) Set(a *baseFieldElem) *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = a[0], a[1], a[2], a[3], a[4]
	return c
}

// SetBytes uncompresses the little-endian representation of in so that it is
// suitable for arithmetic operations and sets it to c, returning c.
func (c *baseFieldElem) SetBytes(in []byte) *baseFieldElem {
	if len(in) != 16 {
		return nil
	}

	c.SetZero()
	for i := 0; i < 2; i++ {
		for j := 0; j < 26; j++ {
			k := (26+25)*i + j
			c[2*i+0] += getBit(in, k) << uint(j)
		}
		for j := 0; j < 25; j++ {
			k := (26+25)*i + j + 26
			c[2*i+1] += getBit(in, k) << uint(j)
		}
	}
	for j := 0; j < 25; j++ {
		k := j + 102
		c[4] += getBit(in, k) << uint(j)
	}

	return c
}

// SetZero sets c to zero and returns c.
func (c *baseFieldElem) SetZero() *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = 0, 0, 0, 0, 0
	return c
}

// SetOne sets c to 1 and returns c.
func (c *baseFieldElem) SetOne() *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = 1, 0, 0, 0, 0
	return c
}

// IsZero returns true if c is zero.
func (c *baseFieldElem) IsZero() bool {
	return c[0] == 0 && c[1] == 0 && c[2] == 0 && c[3] == 0 && c[4] == 0
}

// Neg sets c to be -a mod p, and returns c.
func (c *baseFieldElem) Neg(a *baseFieldElem) *baseFieldElem {
	c[0] = bNeg(a[0])
	c[1] = aNeg(a[1])
	c[2] = bNeg(a[2])
	c[3] = aNeg(a[3])
	c[4] = aNeg(a[4])

	c.reduce()
	return c
}

// bfeDbl sets c to be 2*a.
func bfeDbl(c, a *baseFieldElem) {
	bfeAdd(c, a, a)
}

// bfeAdd sets c to be a+b mod p.
func bfeAdd(c, a, b *baseFieldElem) {
	// The first carry is manually inlined for performance.
	var carry uint64

	c[0], carry = bSplit(a[0] + b[0])
	c[1], carry = aSplit(a[1] + b[1] + carry)
	c[2], carry = bSplit(a[2] + b[2] + carry)
	c[3], carry = aSplit(a[3] + b[3] + carry)
	c[4], carry = aSplit(a[4] + b[4] + carry)
	c[0] += carry

	c.carry()
	c.reduce()
}

// bfeSub sets c to be a-b mod p.
func bfeSub(c, a, b *baseFieldElem) {
	var carry uint64

	c[0], carry = bSplit(a[0] + bNeg(b[0]))
	c[1], carry = aSplit(a[1] + aNeg(b[1]) + carry)
	c[2], carry = bSplit(a[2] + bNeg(b[2]) + carry)
	c[3], carry = aSplit(a[3] + aNeg(b[3]) + carry)
	c[4], carry = aSplit(a[4] + aNeg(b[4]) + carry)
	c[0] += carry

	c.reduce()
}

// bfeMul sets c to be a*b mod p.
func bfeMul(c, a, b *baseFieldElem) {
	var (
		l0, m0, u0 = baSplit(a[0]*b[0] + 2*(a[4]*b[1]+a[3]*b[2]+a[2]*b[3]+a[1]*b[4]))
		l1, m1, u1 = abSplit(a[1]*b[0] + a[0]*b[1] + a[4]*b[2] + a[2]*b[4] + 2*(a[3]*b[3]))
		l2, m2, u2 = baSplit(a[2]*b[0] + a[0]*b[2] + 2*(a[4]*b[3]+a[3]*b[4]+a[1]*b[1]))
		l3, m3, u3 = aaSplit(a[3]*b[0] + a[2]*b[1] + a[1]*b[2] + a[0]*b[3] + a[4]*b[4])
		l4, m4, u4 = abSplit(a[4]*b[0] + a[2]*b[2] + a[0]*b[4] + 2*(a[3]*b[1]+a[1]*b[3]))
	)

	var carry uint64

	c[0], carry = bSplit(l0 + m4 + u3)
	c[1], carry = aSplit(l1 + m0 + u4 + carry)
	c[2], carry = bSplit(l2 + m1 + u0 + carry)
	c[3], carry = aSplit(l3 + m2 + u1 + carry)
	c[4], carry = aSplit(l4 + m3 + u2 + carry)
	c[0] += carry

	c.carry()
	c.reduce()
}

// bfeSquare sets c to a^2 mod p.
func bfeSquare(c, a *baseFieldElem) {
	var (
		l0, m0, u0 = baSplit(a[0]*a[0] + 4*(a[2]*a[3]+a[1]*a[4]))
		l1, m1, u1 = abSplit(2 * (a[0]*a[1] + a[2]*a[4] + a[3]*a[3]))
		l2, m2, u2 = baSplit(2 * (a[0]*a[2] + a[1]*a[1] + 2*a[3]*a[4]))
		l3, m3, u3 = aaSplit(a[4]*a[4] + 2*(a[1]*a[2]+a[0]*a[3]))
		l4, m4, u4 = abSplit(a[2]*a[2] + 2*(a[0]*a[4]+2*a[1]*a[3]))
	)

	var carry uint64

	c[0], carry = bSplit(l0 + m4 + u3)
	c[1], carry = aSplit(l1 + m0 + u4 + carry)
	c[2], carry = bSplit(l2 + m1 + u0 + carry)
	c[3], carry = aSplit(l3 + m2 + u1 + carry)
	c[4], carry = aSplit(l4 + m3 + u2 + carry)
	c[0] += carry

	c.carry()
	c.reduce()
}

// TODO(brendan): Delete me.
func (c *baseFieldElem) Exp(a *baseFieldElem, power *big.Int) *baseFieldElem {
	sum := newBaseFieldElem().SetOne()
	t := newBaseFieldElem()

	for i := power.BitLen() - 1; i >= 0; i-- {
		bfeSquare(t, sum)
		if power.Bit(i) != 0 {
			bfeMul(sum, t, a)
		} else {
			sum.Set(t)
		}
	}

	c.Set(sum)
	return c
}

// TODO(brendan): Make me efficient.
func (c *baseFieldElem) Invert(a *baseFieldElem) *baseFieldElem {
	power := big.NewInt(1)
	power.Lsh(power, 127).Sub(power, big.NewInt(3))

	return c.Exp(a, power)
}

// carry pushes excess data in the upper bits of each word into the next word,
// reducing mod p if necessary.
func (c *baseFieldElem) carry() {
	var carry uint64

	c[0], carry = bSplit(c[0])
	c[1] += carry
	c[1], carry = aSplit(c[1])
	c[2] += carry
	c[2], carry = bSplit(c[2])
	c[3] += carry
	c[3], carry = aSplit(c[3])
	c[4] += carry
	c[4], carry = aSplit(c[4])
	c[0] += carry
}

// reduce will set c to zero if it is equal to p. This is the only case where c
// will not naturally be reduced to canonical form by c.carry().
func (c *baseFieldElem) reduce() {
	hamming := c[0] + c[1] + c[2] + c[3] + c[4]
	if hamming == 0x0dfffffb {
		c.SetZero()
	}
}
