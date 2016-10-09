package fourq

import (
	"math/big"
)

var (
	Order, _ = new(big.Int).SetString("73846995687063900142583536357581573884798075859800097461294096333596429543", 10)
	Gx, Gy   *big.Int

	// Constants exclusively for tests.
	p, _ = new(big.Int).SetString("170141183460469231731687303715884105727", 10)
)
