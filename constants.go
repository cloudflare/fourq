package fourq

import (
	"math/big"
)

var (
	Order, _ = new(big.Int).SetString("73846995687063900142583536357581573884798075859800097461294096333596429543", 10)

	aMask uint64 = 0x7fffffffffffffff
	bMask uint64 = 0xffffffffffffffff

	g = &point{
		x: fieldElem{
			x: baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305},
			y: baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c},
		},
		y: fieldElem{
			x: baseFieldElem{0xb924a2462bcbb287, 0x0e3fee9ba120785a},
			y: baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242},
		},
		t: fieldElem{
			x: baseFieldElem{0x894ba36ee8cee416, 0x35bfa1947fb0913e},
			y: baseFieldElem{0x673c574d296cd8d0, 0x7bfb41a38e7076ac},
		},
		z: fieldElem{
			x: baseFieldElem{0x0000000000000001, 0x0000000000000000},
			y: baseFieldElem{0x0000000000000000, 0x0000000000000000},
		},
	}
	Gx, _ = new(big.Int).SetString("aa33387bad92652805b32f7c2372341af677ac60b39f86969caa78283f551f1e", 16)
	Gy, _ = new(big.Int).SetString("87b2cb2b46a224b95a7820a19bee3f0e5c8b4c8444c3a74942020e63f84a1c6e", 16)

	d = &fieldElem{
		x: baseFieldElem{0x142, 0xe4},
		y: baseFieldElem{0xb3821488f1fc0c8d, 0x5e472f846657e0fc},
	}
	one = &fieldElem{x: baseFieldElem{0x1, 0x0}}
)
