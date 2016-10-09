// +build amd64,!noasm

package fourq

import (
	"fmt"
	"math/big"
)

var (
	aMask uint64 = 0x7fffffffffffffff
	bMask uint64 = 0xffffffffffffffff

	g = &point{
		x: &gfP2{
			x: &baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305},
			y: &baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c},
		},
		y: &gfP2{
			x: &baseFieldElem{0xb924a2462bcbb287, 0xe3fee9ba120785a},
			y: &baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242},
		},
		z: newGFp2().SetOne(),
	}

	d = &gfP2{
		x: &baseFieldElem{0x142, 0xe4},
		y: &baseFieldElem{0xb3821488f1fc0c8d, 0x5e472f846657e0fc},
	}
)

func init() {
	g.t = newGFp2().Mul(g.x, g.y)
	Gx, Gy = g.Int()
}

func getBit(in []byte, k int) uint64 {
	return uint64(in[k/8]>>uint(k%8)) & 1
}

func aNeg(in uint64) uint64 { return (^in) & aMask }
func bNeg(in uint64) uint64 { return (^in) & bMask }

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [2]uint64

func newBaseFieldElem() *baseFieldElem {
	return &baseFieldElem{}
}

// numToBFE takes a big.Int as input and returns its representation as a
// baseFieldElement. This function is used exclusively in tests.
func numToBFE(in *big.Int) *baseFieldElem {
	out := newBaseFieldElem()

	for i := 0; i < 64; i++ {
		out[0] += uint64(in.Bit(i)) << uint(i)
	}
	for i := 0; i < 64; i++ {
		out[1] += uint64(in.Bit(i+64)) << uint(i)
	}

	return out
}

func (c *baseFieldElem) String() string {
	return fmt.Sprintf("%16.16x%16.16x", c[1], c[0])
}

// Bytes returns the compressed, little-endian representation of a number.
// It is compatible with c.SetBytes().
func (c *baseFieldElem) Bytes() []byte {
	out := make([]byte, 16)

	for i := uint(0); i < 64; i++ {
		b := byte(c[0]>>i) & 1
		out[i/8] += b << (i % 8)
	}
	for i := uint(64); i < 128; i++ {
		b := byte(c[1]>>(i-64)) & 1
		out[i/8] += b << (i % 8)
	}

	return out
}

// Set sets c to be a (with duplication) and returns c.
func (c *baseFieldElem) Set(a *baseFieldElem) *baseFieldElem {
	c[0], c[1] = a[0], a[1]
	return c
}

// SetBytes uncompresses the little-endian representation of in so that it is
// suitable for arithmetic operations and sets it to c, returning c.
func (c *baseFieldElem) SetBytes(in []byte) *baseFieldElem {
	if len(in) != 16 {
		return nil
	}

	c.SetZero()
	for i := 0; i < 64; i++ {
		c[0] += getBit(in, i) << uint(i)
	}
	for i := 64; i < 128; i++ {
		c[1] += getBit(in, i) << uint(i-64)
	}

	return c
}

// SetZero sets c to zero and returns c.
func (c *baseFieldElem) SetZero() *baseFieldElem {
	c[0], c[1] = 0, 0
	return c
}

// SetOne sets c to 1 and returns c.
func (c *baseFieldElem) SetOne() *baseFieldElem {
	c[0], c[1] = 1, 0
	return c
}

// IsZero returns true if c is zero.
func (c *baseFieldElem) IsZero() bool {
	return c[0] == 0 && c[1] == 0
}

// Neg sets c to be -a mod p, and returns c.
func (c *baseFieldElem) Neg(a *baseFieldElem) *baseFieldElem {
	c[0] = bNeg(a[0])
	c[1] = aNeg(a[1])

	c.reduce()
	return c
}

// bfeDbl sets c to be 2*a.
func bfeDbl(c, a *baseFieldElem) {
	bfeAdd(c, a, a)
}

//go:noescape

// bfeAdd sets c to be a+b mod p.
func bfeAdd(c, a, b *baseFieldElem)

// bfeSub sets c to be a-b mod p.
func bfeSub(c, a, b *baseFieldElem) {
	bN := newBaseFieldElem()
	bN.Neg(b)
	bfeAdd(c, a, bN)
}

//go:noescape

// bfeMul sets c to be a*b mod p.
func bfeMul(c, a, b *baseFieldElem)

//go:noescape

// bfeSquare sets c to a^2 mod p.
func bfeSquare(c, a *baseFieldElem)

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

// reduce will set c to zero if it is equal to p. This is the only case where c
// will not naturally be reduced to canonical form by c.carry().
func (c *baseFieldElem) reduce() {
	if c[0] == bMask && c[1] == aMask {
		c.SetZero()
	}
}
