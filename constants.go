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
	Gx, _ = new(big.Int).SetString("0e3fee9ba120785ab924a2462bcbb287", 16)
	Gy, _ = new(big.Int).SetString("6e1c4af8630e024249a7c344844c8b5c", 16)

	d = &fieldElem{
		x: baseFieldElem{0x142, 0xe4},
		y: baseFieldElem{0xb3821488f1fc0c8d, 0x5e472f846657e0fc},
	}
	one = &fieldElem{x: baseFieldElem{0x1, 0x0}}

	// generatorBase contains pre-computed multiples of the curve's generator,
	// to speed up ScalarBaseMult.
	generatorBase = [...]*point{
		nil,
		&point{ // [1]G
			x: fieldElem{x: baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305}, y: baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c}},
			y: fieldElem{x: baseFieldElem{0xb924a2462bcbb287, 0x0e3fee9ba120785a}, y: baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242}},
			z: fieldElem{x: baseFieldElem{0x0000000000000001, 0x0000000000000000}, y: baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			t: fieldElem{x: baseFieldElem{0x14bd7e65d5ed215d, 0x6ca4689baab64be3}, y: baseFieldElem{0xd40c4d1c99983426, 0x065795b9051a0f93}},
		},
		&point{ // [2^128]G
			x: fieldElem{x: baseFieldElem{0x70967161c37ef3ea, 0x577055744cfa93e5}, y: baseFieldElem{0xee485cf0e5311667, 0x16c9ec3138054186}},
			y: fieldElem{x: baseFieldElem{0x55f77a6f2d1337a5, 0x5f3d60575044c1ed}, y: baseFieldElem{0x485c522f6deb3768, 0x25f3c70241d31ee4}},
			z: fieldElem{x: baseFieldElem{0x0000000000000001, 0x0000000000000000}, y: baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			t: fieldElem{x: baseFieldElem{0x72d91df80a5c939b, 0x7c697b4f3849639b}, y: baseFieldElem{0x0b8c099f1e9f70fc, 0x359c95714b932845}},
		},
		&point{ // [2^128 + 1]G
			x: fieldElem{x: baseFieldElem{0xd42a47f8a1c12c67, 0x3a5d6736fd87c598}, y: baseFieldElem{0x3eb11565a0fb1722, 0x1e2996ff1802978f}},
			y: fieldElem{x: baseFieldElem{0xa37999ba33ab621e, 0x22ff4a6b711e0583}, y: baseFieldElem{0x3118f0d076d605d3, 0x69c94e16f3c9ec8e}},
			z: fieldElem{x: baseFieldElem{0x0000000000000001, 0x0000000000000000}, y: baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			t: fieldElem{x: baseFieldElem{0xfef46e2e1406c46c, 0x06096ee3b8e7e48f}, y: baseFieldElem{0x2977a7b5f5ec7b3d, 0x1ac274ae035666d3}},
		},
	}
)
