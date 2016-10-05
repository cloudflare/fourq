package fourq

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

func aNeg(in uint64) uint64 { return (^in) & aMask }
func bNeg(in uint64) uint64 { return (^in) & bMask }

// baseFieldElem is an element of the curve's base field, the integers modulo
// p=2^127-1. baseFieldElem is always in reduced form.
type baseFieldElem [5]uint64

func (c *baseFieldElem) String() string {
	return fmt.Sprintf("%x %x %x %x %x", c[0], c[1], c[2], c[3], c[4])
}

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

func (c *baseFieldElem) Set(a *baseFieldElem) *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = a[0], a[1], a[2], a[3], a[4]
	return c
}

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

func (c *baseFieldElem) SetZero() *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = 0, 0, 0, 0, 0
	return c
}

func (c *baseFieldElem) SetOne() *baseFieldElem {
	c[0], c[1], c[2], c[3], c[4] = 1, 0, 0, 0, 0
	return c
}

func (c *baseFieldElem) IsZero() bool {
	return c[0] == 0 && c[1] == 0 && c[2] == 0 && c[3] == 0 && c[4] == 0
}

func (c *baseFieldElem) Neg(a *baseFieldElem) *baseFieldElem {
	c[0] = bNeg(a[0])
	c[1] = aNeg(a[1])
	c[2] = bNeg(a[2])
	c[3] = aNeg(a[3])
	c[4] = aNeg(a[4])

	c.reduce()
	return c
}

func (c *baseFieldElem) Dbl(a *baseFieldElem) *baseFieldElem {
	return c.Add(a, a)
}

func (c *baseFieldElem) Add(a, b *baseFieldElem) *baseFieldElem {
	c[0] = a[0] + b[0]
	c[1] = a[1] + b[1]
	c[2] = a[2] + b[2]
	c[3] = a[3] + b[3]
	c[4] = a[4] + b[4]

	c.carry()
	c.reduce()
	return c
}

func (c *baseFieldElem) Sub(a, b *baseFieldElem) *baseFieldElem {
	c[0] = a[0] + bNeg(b[0])
	c[1] = a[1] + aNeg(b[1])
	c[2] = a[2] + bNeg(b[2])
	c[3] = a[3] + aNeg(b[3])
	c[4] = a[4] + aNeg(b[4])

	c.carry()
	c.reduce()
	return c
}

func (c *baseFieldElem) Mul(a, b *baseFieldElem) *baseFieldElem {
	var (
		l0, u0 = bSplit(a[0]*b[0] + 2*(a[4]*b[1]+a[3]*b[2]+a[2]*b[3]+a[1]*b[4]))
		l1, u1 = aSplit(a[1]*b[0] + a[0]*b[1] + a[4]*b[2] + a[2]*b[4] + 2*(a[3]*b[3]))
		l2, u2 = bSplit(a[2]*b[0] + a[0]*b[2] + 2*(a[4]*b[3]+a[3]*b[4]+a[1]*b[1]))
		l3, u3 = aSplit(a[3]*b[0] + a[2]*b[1] + a[1]*b[2] + a[0]*b[3] + a[4]*b[4])
		l4, u4 = aSplit(a[4]*b[0] + a[2]*b[2] + a[0]*b[4] + 2*(a[3]*b[1]+a[1]*b[3]))
	)

	c[0] = l0 + u4
	c[1] = l1 + u0
	c[2] = l2 + u1
	c[3] = l3 + u2
	c[4] = l4 + u3

	c.carry()
	c.carry()
	c.reduce()
	return c
}

func (c *baseFieldElem) Square(a *baseFieldElem) *baseFieldElem {
	var (
		l0, u0 = bSplit(a[0]*a[0] + 4*(a[2]*a[3]+a[1]*a[4]))
		l1, u1 = aSplit(2 * (a[0]*a[1] + a[2]*a[4] + a[3]*a[3]))
		l2, u2 = bSplit(2 * (2*a[3]*a[4] + a[0]*a[2] + a[1]*a[1]))
		l3, u3 = aSplit(a[4]*a[4] + 2*(a[1]*a[2]+a[0]*a[3]))
		l4, u4 = aSplit(a[2]*a[2] + 2*(2*a[1]*a[3]+a[0]*a[4]))
	)

	c[0] = l0 + u4
	c[1] = l1 + u0
	c[2] = l2 + u1
	c[3] = l3 + u2
	c[4] = l4 + u3

	c.carry()
	c.carry()
	c.reduce()
	return c
}

// TODO(brendan): Delete me.
func (c *baseFieldElem) Exp(a *baseFieldElem, power *big.Int) *baseFieldElem {
	sum := (&baseFieldElem{}).SetOne()
	t := &baseFieldElem{}

	for i := power.BitLen() - 1; i >= 0; i-- {
		t.Square(sum)
		if power.Bit(i) != 0 {
			sum.Mul(t, a)
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
	zero := true
	zero = zero && c[0] == bMask
	zero = zero && c[1] == aMask
	zero = zero && c[2] == bMask
	zero = zero && c[3] == aMask
	zero = zero && c[4] == aMask

	if zero {
		c.SetZero()
	}
}
