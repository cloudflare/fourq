package fourq

import (
	"fmt"
	"math/big"
)

func getBit(in []byte, k int) uint64 {
	return uint64(in[k/8]>>uint(k%8)) & 1
}

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

// reduce will set c to zero if it is equal to p. This is the only case where c
// will not naturally be reduced to canonical form by c.carry().
func (c *baseFieldElem) reduce() {
	if c[0] == bMask && c[1] == aMask {
		c.SetZero()
	}
}
