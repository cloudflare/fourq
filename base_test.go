package fourq

import (
	"testing"

	"crypto/rand"
	"fmt"
	"math/big"
)

type intFunc1 func(a *big.Int) *big.Int
type elemFunc1 func(C, A *baseFieldElem)

func randomTest1(t *testing.T, real intFunc1, cand elemFunc1) {
	for i := 0; i < 10000; i++ {
		a, _ := rand.Int(rand.Reader, p)
		c := real(a)
		c.Mod(c, p)

		A := numToBFE(a)
		C := newBaseFieldElem()
		cand(C, A)

		if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
			t.Log(a)
			t.Log(C)
			t.Fatal("Incorrect output.")
		}
	}
}

type intFunc2 func(a, b *big.Int) *big.Int
type elemFunc2 func(C, A, B *baseFieldElem)

func randomTest2(t *testing.T, real intFunc2, cand elemFunc2) {
	for i := 0; i < 10000; i++ {
		a, _ := rand.Int(rand.Reader, p)
		b, _ := rand.Int(rand.Reader, p)
		c := real(a, b)
		c.Mod(c, p)

		A, B := numToBFE(a), numToBFE(b)
		C := newBaseFieldElem()
		cand(C, A, B)

		if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
			t.Log(a, b)
			t.Log(c, C)
			t.Fatal("Incorrect output.")
		}
	}
}

// The following are stochastic tests, to hopefully find broken edge-cases that
// wouldn't be explicitly tested for.

func TestBaseAdd(t *testing.T) { randomTest2(t, new(big.Int).Add, bfeAdd) }
func TestBaseSub(t *testing.T) { randomTest2(t, new(big.Int).Sub, bfeSub) }
func TestBaseMul(t *testing.T) { randomTest2(t, new(big.Int).Mul, bfeMul) }

//
// func TestBaseNeg(t *testing.T) { randomTest1(t, new(big.Int).Neg, new(baseFieldElem).Neg) }

func TestBaseSquare(t *testing.T) {
	square := func(a *big.Int) *big.Int { return new(big.Int).Mul(a, a) }
	randomTest1(t, square, bfeSquare)
}

// func TestBaseInvert(t *testing.T) {
// 	invert := func(a *big.Int) *big.Int { return new(big.Int).ModInverse(a, p) }
// 	randomTest1(t, invert, new(baseFieldElem).Invert)
// }

// TestBaseAddFull adds -1 to -1. It should be the worst-case for carries that
// baseFieldElem.Add sees.
func TestBaseAddFull(t *testing.T) {
	one := big.NewInt(1)
	a := new(big.Int).Lsh(one, 127)
	a.Sub(a, one).Sub(a, one)
	c := new(big.Int).Add(a, a)
	c.Mod(c, p)

	A, C := numToBFE(a), newBaseFieldElem()
	bfeAdd(C, A, A)

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

	A, B, C := numToBFE(a), numToBFE(b), newBaseFieldElem()
	bfeAdd(C, A, B)
	C.reduce()

	if fmt.Sprint(numToBFE(c)) != fmt.Sprint(C) {
		t.Log(a, b)
		t.Log(C)
		t.Fatal("Incorrect output.")
	}
}

// TestMulZero multiplies zero by one and checks that zero is the result.
func TestBaseMulZero(t *testing.T) {
	a, b := big.NewInt(0), big.NewInt(1)

	A, B, C := numToBFE(a), numToBFE(b), newBaseFieldElem()
	bfeMul(C, A, B)

	if fmt.Sprint(numToBFE(a)) != fmt.Sprint(C) {
		t.Fatalf("Incorrect output: %v", C)
	}
}

func BenchmarkBaseAdd(b *testing.B) {
	A, B, C := newBaseFieldElem(), newBaseFieldElem(), newBaseFieldElem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bfeAdd(C, A, B)
	}
}

func BenchmarkBaseMul(b *testing.B) {
	A, B, C := newBaseFieldElem(), newBaseFieldElem(), newBaseFieldElem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bfeMul(C, A, B)
	}
}

func BenchmarkBaseSquare(b *testing.B) {
	A, C := newBaseFieldElem(), newBaseFieldElem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bfeSquare(C, A)
	}
}
