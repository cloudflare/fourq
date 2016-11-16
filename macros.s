#define bfeReduce(c0,c1) \
	BTRQ $63, c1 \
	ADCQ $0, c0 \
	ADCQ $0, c1

#define bfeNeg(c0,c1) \
	NOTQ c0 \
	NOTQ c1 \
	BTRQ $63, c1

#define bfeDbl(c0,c1) \
	SHLQ $1, c1 \
	SHLQ $1, c0 \
	ADCQ $0, c1 \
	bfeReduce(c0,c1)

#define bfeAdd(a0,a1, c0,c1) \
	ADDQ a0, c0 \
	ADCQ a1, c1 \
	bfeReduce(c0,c1)

#define bfeSub(a0,a1, c0,c1) \
	bfeNeg(c0,c1) \
	bfeAdd(a0,a1, c0,c1)

#define bfeReverseSub(a0,a1, c0,c1) \
	bfeNeg(c0,c1) \
	bfeAdd(c0,c1, a0,a1)

#define bfeMul(carry, a0,a1, b0,b1, c0,c1) \
	MOVQ $0, carry \
	\
	MOVQ a0, AX \
	MULQ b0 \
	MOVQ AX, c0 \
	MOVQ DX, c1 \
	\
	bfeMulCore(carry, a0,a1, b0,b1, c0, c1)

#define bfeMulAdd(carry, a0,a1, b0,b1, c0,c1) \
	MOVQ a0, AX \
	MULQ b0 \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry \
	\
	bfeMulCore(carry, a0,a1, b0,b1, c0, c1)

#define bfeMulCore(carry, a0,a1, b0,b1, c0,c1) \
	MOVQ a0, AX \
	MULQ b1 \
	SHLQ $1, DX \
	ADDQ DX, c0 \
	ADCQ AX, c1 \
	ADCQ $0, carry \
	\
	MOVQ a1, AX \
	MULQ b0 \
	SHLQ $1, DX \
	ADDQ DX, c0 \
	ADCQ AX, c1 \
	ADCQ $0, carry \
	\
	MOVQ a1, AX \
	MULQ b1 \
	SHLQ $1, DX \
	SHLQ $1, AX \
	ADCQ $0, DX \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry

#define bfeMulReduce(carry, c0,c1) \
	SHLQ $1, carry \
	BTRQ $63, c1 \
	ADCQ carry, c0 \
	ADCQ $0, c1 \
	\
	bfeReduce(c0,c1)

#define bfeSquare(carry, a0,a1, c0,c1) \
	MOVQ $0, carry \
	\
	MOVQ a0, AX \
	MULQ a0 \
	MOVQ AX, c0 \
	MOVQ DX, c1 \
	\
	bfeSquareCore(carry, a0,a1, c0,c1)

#define bfeSquareAdd(carry, a0,a1, c0,c1) \
	MOVQ a0, AX \
	MULQ a0 \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry \
	\
	bfeSquareCore(carry, a0,a1, c0,c1)

#define bfeSquareCore(carry, a0,a1, c0,c1) \
	MOVQ a0, AX \
	MULQ a1 \
	SHLQ $1, DX \
	ADDQ DX, c0 \
	ADCQ AX, c1 \
	ADCQ $0, carry \
	ADDQ DX, c0 \
	ADCQ AX, c1 \
	ADCQ $0, carry \
	\
	MOVQ a1, AX \
	MULQ a1 \
	SHLQ $1, DX \
	SHLQ $1, AX \
	ADCQ $0, DX \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry

#define feMov(a0,a1,a2,a3, c0,c1,c2,c3) \
	MOVQ a0, c0 \
	MOVQ a1, c1 \
	MOVQ a2, c2 \
	MOVQ a3, c3

#define feDbl(c0,c1,c2,c3) \
	bfeDbl(c0,c1) \
	bfeDbl(c2,c3)

#define feAdd(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeAdd(a0,a1, c0,c1) \
	bfeAdd(a2,a3, c2,c3)

#define feSub(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeSub(a0,a1, c0,c1) \
	bfeSub(a2,a3, c2,c3)

#define feReverseSub(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeReverseSub(a0,a1, c0,c1) \
	bfeReverseSub(a2,a3, c2,c3)

#define feMul(carry, a0,a1,a2,a3, b0,b1,b2,b3, c0,c1,c2,c3) \
	bfeMul(carry, a2,a3, b2,b3, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	bfeNeg(c0,c1) \
	MOVQ $0, carry \
	bfeMulAdd(carry, a0,a1, b0,b1, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	\
	bfeMul(carry, a0,a1, b2,b3, c2,c3) \
	bfeMulAdd(CX, a2,a3, b0,b1, c2,c3) \
	bfeMulReduce(carry, c2,c3)

#define feSquare(carry,t0,t1, a0,a1,a2,a3, c0,c1,c2,c3) \
	MOVQ a0, c2 \
	MOVQ a1, c3 \
	bfeAdd(a2,a3, c2,c3) \
	\
	MOVQ a2, t0 \
	MOVQ a3, t1 \
	bfeSub(a0,a1, t0,t1) \
	\
	bfeMul(carry, c2,c3, t0,t1, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	\
	bfeMul(carry, a0,a1, a2,a3, c2,c3) \
	bfeMulReduce(carry, c2,c3) \
	bfeDbl(c2,c3)
