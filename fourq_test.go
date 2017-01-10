package fourq

import (
	"testing"

	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/curve25519"
)

func TestIsOnCurve(t *testing.T) {
	if !IsOnCurve(G) {
		t.Fatal("Generator is not on curve.")
	}

	pt2 := ScalarMult(G, Order.Bytes())
	if !IsOnCurve(pt2) {
		t.Fatal("Identity point is not on curve.")
	}

	k := make([]byte, 32)
	rand.Read(k)
	pt3 := ScalarMult(G, k)
	if !IsOnCurve(pt3) {
		t.Fatal("Random multiple of generator is not on curve.")
	}

	pt4 := [32]byte{}
	pt4[0], pt4[16] = 5, 7
	if IsOnCurve(pt4) {
		t.Fatal("Non-existent point is on curve.")
	}
}

func TestScalarBaseMultOrder(t *testing.T) {
	pt3 := ScalarBaseMult(Order.Bytes())
	if pt3 != [32]byte{1} {
		t.Fatal("ScalarMult(Generator, Order) was not identity.")
	}

	k := make([]byte, 32)
	rand.Read(k)
	pt4, pt5 := ScalarMult(G, k), ScalarBaseMult(k)
	if pt4 != pt5 {
		t.Fatal("ScalarMult(G, k) != ScalarBaseMult(k)")
	}
}

func TestScalarMult(t *testing.T) {
	// Source: https://github.com/bifurcation/fourq/blob/master/impl/curve4q.py#L549
	scalar := [4]uint64{0x3AD457AB55456230, 0x3A8B3C2C6FD86E0C, 0x7E38F7C9CFBB9166, 0x0028FD6CBDA458F0}

	pt := G
	for i := 0; i < 1000; i++ {
		scalar[1] = scalar[2]
		scalar[2] += scalar[0]
		scalar[2] &= 0xffffffffffffffff

		k := new(big.Int).SetUint64(scalar[3])
		k.Lsh(k, 64)
		k.Add(k, new(big.Int).SetUint64(scalar[2]))
		k.Lsh(k, 64)
		k.Add(k, new(big.Int).SetUint64(scalar[1]))
		k.Lsh(k, 64)
		k.Add(k, new(big.Int).SetUint64(scalar[0]))

		pt = ScalarMult(pt, k.Bytes())
	}

	real := "44336f9967501c286c930e7c81b3010945125f9129c4e84f10e2acac8e940b57"
	if fmt.Sprintf("%x", pt) != real {
		t.Fatal("Point is wrong!")
	}
}

func BenchmarkScalarBaseMult(b *testing.B) {
	k := make([]byte, 32)
	rand.Read(k)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ScalarBaseMult(k)
	}
}

func BenchmarkScalarMult(b *testing.B) {
	k := make([]byte, 32)
	rand.Read(k)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ScalarMult(G, k)
	}
}

func BenchmarkP256Base(b *testing.B) {
	c := elliptic.P256()

	k := make([]byte, 32)
	rand.Read(k)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ScalarBaseMult(k)
	}
}

func BenchmarkP256(b *testing.B) {
	c := elliptic.P256()
	params := c.Params()

	k := make([]byte, 32)
	rand.Read(k)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ScalarMult(params.Gx, params.Gy, k)
	}
}

func BenchmarkCurve25519(b *testing.B) {
	dst, in := [32]byte{}, [32]byte{}
	rand.Read(in[:])

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve25519.ScalarBaseMult(&dst, &in)
	}
}
