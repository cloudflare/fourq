// +build amd64,!noasm

package fourq

import (
	"math/big"
)

var (
	Order, _ = new(big.Int).SetString("73846995687063900142583536357581573884798075859800097461294096333596429543", 10)
	Gx, Gy   *big.Int

	aMask uint64 = 0x7fffffffffffffff
	bMask uint64 = 0xffffffffffffffff

	g = &point{
		x: gfP2{
			x: baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305},
			y: baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c},
		},
		y: gfP2{
			x: baseFieldElem{0xb924a2462bcbb287, 0xe3fee9ba120785a},
			y: baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242},
		},
		z: *newGFp2().SetOne(),
	}

	d = &gfP2{
		x: baseFieldElem{0x142, 0xe4},
		y: baseFieldElem{0xb3821488f1fc0c8d, 0x5e472f846657e0fc},
	}

	// Constants exclusively for tests.
	p, _ = new(big.Int).SetString("170141183460469231731687303715884105727", 10)
)

func init() {
	feMul(&g.t, &g.x, &g.y)
	Gx, Gy = g.Int()
}
