package fourq

import (
	"testing"

	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	p, _ = new(big.Int).SetString("170141183460469231731687303715884105727", 10)
	one  = big.NewInt(1)
)

// numToBFE takes a big.Int as input and returns its representation as a
// baseFieldElement.
func numToBFE(in *big.Int) *baseFieldElem {
	out := &baseFieldElem{}

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

type intFunc1 func(a *big.Int) *big.Int
type elemFunc1 func(A *baseFieldElem) *baseFieldElem

func randomTest1(t *testing.T, real intFunc1, cand elemFunc1) {
	for i := 0; i < 10000; i++ {
		a, _ := rand.Int(rand.Reader, p)
		c := real(a)
		c.Mod(c, p)

		A := numToBFE(a)
		C := cand(A)

		if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
			t.Log(a)
			t.Log(C)
			t.Fatal("Incorrect output.")
		}
	}
}

type intFunc2 func(a, b *big.Int) *big.Int
type elemFunc2 func(A, B *baseFieldElem) *baseFieldElem

func randomTest2(t *testing.T, real intFunc2, cand elemFunc2) {
	for i := 0; i < 10000; i++ {
		a, _ := rand.Int(rand.Reader, p)
		b, _ := rand.Int(rand.Reader, p)
		c := real(a, b)
		c.Mod(c, p)

		A, B := numToBFE(a), numToBFE(b)
		C := cand(A, B)

		if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
			t.Log(a, b)
			t.Log(c, C)
			t.Fatal("Incorrect output.")
		}
	}
}

// The following are stochastic tests, to hopefully find broken edge-cases that
// wouldn't be explicitly tested for.

func TestBaseAdd(t *testing.T) { randomTest2(t, new(big.Int).Add, new(baseFieldElem).Add) }
func TestBaseSub(t *testing.T) { randomTest2(t, new(big.Int).Sub, new(baseFieldElem).Sub) }
func TestBaseMul(t *testing.T) { randomTest2(t, new(big.Int).Mul, new(baseFieldElem).Mul) }

func TestBaseNeg(t *testing.T) { randomTest1(t, new(big.Int).Neg, new(baseFieldElem).Neg) }

func TestBaseSquare(t *testing.T) {
	square := func(a *big.Int) *big.Int { return new(big.Int).Mul(a, a) }
	randomTest1(t, square, new(baseFieldElem).Square)
}

func TestBaseInvert(t *testing.T) {
	invert := func(a *big.Int) *big.Int { return new(big.Int).ModInverse(a, p) }
	randomTest1(t, invert, new(baseFieldElem).Invert)
}

// TestBaseAddFull adds -1 to -1. It should be the worst-case for carries that
// baseFieldElem.Add sees.
func TestBaseAddFull(t *testing.T) {
	a := big.NewInt(1)
	a.Lsh(a, 127).Sub(a, one).Sub(a, one)
	c := new(big.Int).Add(a, a)
	c.Mod(c, p)

	A, C := numToBFE(a), &baseFieldElem{}
	C.Add(A, A)

	if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
		t.Fatalf("Incorrect output: %v", C)
	}
}

// TEstBaseAddNegatives adds a random number to its negative. This checks that
// carries as well as the final reduction are happening properly.
func TestBaseAddNegatives(t *testing.T) {
	a, _ := rand.Int(rand.Reader, p)
	b := new(big.Int).Neg(a)
	b.Mod(b, p)
	c := big.NewInt(0)

	A, B, C := numToBFE(a), numToBFE(b), &baseFieldElem{}
	C.Add(A, B)

	if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
		t.Log(a, b)
		t.Log(C)
		t.Fatal("Incorrect output.")
	}
}

// TestMulZero multiplies zero by one and checks that zero is the result.
func TestBaseMulZero(t *testing.T) {
	a, b := big.NewInt(0), big.NewInt(1)

	A, B, C := numToBFE(a), numToBFE(b), &baseFieldElem{}
	C.Mul(A, B)

	if fmt.Sprint(numToBFE(a)) != fmt.Sprint(C) {
		t.Fatalf("Incorrect output: %v", C)
	}
}

func BenchmarkBaseAdd(b *testing.B) {
	A, B, C := &baseFieldElem{}, &baseFieldElem{}, &baseFieldElem{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Add(A, B)
	}
}

func BenchmarkBaseMul(b *testing.B) {
	A, B, C := &baseFieldElem{}, &baseFieldElem{}, &baseFieldElem{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Mul(A, B)
	}
}

func BenchmarkBaseSquare(b *testing.B) {
	A, C := &baseFieldElem{}, &baseFieldElem{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Square(A)
	}
}
