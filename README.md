FourQ
-----

FourQ is a high-speed elliptic curve at the 128-bit security level. This package
contains a pure Go and AMD64-specific implementation. It is incomplete -- please
don't use it.

Preliminary benchmarks:
```
BenchmarkScalarBaseMult-4   	   10000	    103370 ns/op	    1088 B/op	      33 allocs/op
BenchmarkP256-4             	   20000	     67504 ns/op	    2592 B/op	      16 allocs/op
BenchmarkCurve25519-4       	   30000	     50072 ns/op	       0 B/op	       0 allocs/op
```
