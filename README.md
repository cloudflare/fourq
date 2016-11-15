FourQ
-----

FourQ is a high-speed elliptic curve at the 128-bit security level. This package
contains only an AMD64-optimized implementation. It is a preliminary
implementation, and is incomplete. In particular, this package does not
implement FourQ's endomorphisms or lattice reduction techniques.

- Documentation: https://godoc.org/github.com/cloudflare/fourq
- For questions about the code, open an issue or email me at: brendan@cloudflare.com

Most recent benchmarks:
```
BenchmarkScalarBaseMult-4   	   50000	     36227 ns/op	     128 B/op	       4 allocs/op
BenchmarkScalarMult-4       	   30000	     50288 ns/op	     128 B/op	       4 allocs/op

BenchmarkP256Base-4         	  100000	     16279 ns/op	     768 B/op	      12 allocs/op
BenchmarkP256-4             	   20000	     65801 ns/op	    2592 B/op	      16 allocs/op

BenchmarkCurve25519-4       	   30000	     48235 ns/op	       0 B/op	       0 allocs/op
```
