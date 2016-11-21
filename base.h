#define bfeMov(a0,a1, c0,c1) \
	MOVQ a0, c0 \
	MOVQ a1, c1

#define bfeReduce(c0,c1) \
	BTRQ $63, c1 \
	ADCQ $0, c0 \
	ADCQ $0, c1

#define bfeNeg(c0,c1) \
	NOTQ c0 \
	NOTQ c1 \
	BTRQ $63, c1

#define bfeHalf(c0,c1) \
	SHLQ $1, c1 \
	SHRQ $1, c1:c0 \
	SHRQ $1, c0:c1 \
	SHRQ $1, c1

#define bfeDbl(c0,c1) \
	SHLQ $1, c1:c0 \
	SHLQ $1, c0:c1 \
	BTRQ $63, c1

// bfeAdd adds a0:a1 to c0:c1.
#define bfeAdd(a0,a1, c0,c1) \
	ADDQ a0, c0 \
	ADCQ a1, c1 \
	bfeReduce(c0,c1)

// bfeSub stores a0:a1 - c0:c1 in c0:c1.
#define bfeSub(a0,a1, c0,c1) \
	bfeNeg(c0,c1) \
	bfeAdd(a0,a1, c0,c1)

// bfeReverseSub negates c0:c1 and stores a0:a1 - c0:c1 in a0:a1.
#define bfeReverseSub(a0,a1, c0,c1) \
	bfeNeg(c0,c1) \
	bfeAdd(c0,c1, a0,a1)

// bfeMulReduce takes the output `c0:c1` and workspace `carry` of a
// bfeMul/bfeSquare and reduces c0:c1 to canonical form.
#define bfeMulReduce(carry, c0,c1) \
	SHLQ $1, carry \
	BTRQ $63, c1 \
	ADCQ carry, c0 \
	ADCQ $0, c1 \
	\
	bfeReduce(c0,c1)

// bfeMul stores a0:a1 * b0:b1 in c0:c1, using carry as workspace.
#define bfeMul(carry, a0,a1, b0,b1, c0,c1) \
	MOVQ $0, carry \
	\
	MOVQ a0, AX \
	MULQ b0 \
	MOVQ AX, c0 \
	MOVQ DX, c1 \
	\
	bfeMulCore(carry, a0,a1, b0,b1, c0, c1)

// bfeMulAdd adds a0:a1 * b0:b1 to c0:c1. carry is the workspace of the previous
// bfeMul/bfeSquare into c0:c1.
#define bfeMulAdd(carry, a0,a1, b0,b1, c0,c1) \
	MOVQ a0, AX \
	MULQ b0 \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry \
	\
	bfeMulCore(carry, a0,a1, b0,b1, c0, c1)

// Not to be called outside of this file.
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

// bfeSquare stores a0:a1^2 in c0:c1, using carry as workspace.
#define bfeSquare(carry, a0,a1, c0,c1) \
	MOVQ $0, carry \
	\
	MOVQ a0, AX \
	MULQ a0 \
	MOVQ AX, c0 \
	MOVQ DX, c1 \
	\
	bfeSquareCore(carry, a0,a1, c0,c1)

// bfeSquareAdd adds a0:a1^2 to c0:c1. carry is the workspace of the previous
// bfeMul/bfeSquare into c0:c1.
#define bfeSquareAdd(carry, a0,a1, c0,c1) \
	MOVQ a0, AX \
	MULQ a0 \
	ADDQ AX, c0 \
	ADCQ DX, c1 \
	ADCQ $0, carry \
	\
	bfeSquareCore(carry, a0,a1, c0,c1)

// Not to be called outside of this file.
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
