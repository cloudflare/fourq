#include "macros.s"

// func feAdd(c, a, b *gfP2)
TEXT ·feAdd(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	feMov(0(DI),8(DI),16(DI),24(DI), AX,BX,CX,DX)
	feAdd(0(SI),8(SI),16(SI),24(SI), AX,BX,CX,DX)

	MOVQ c+0(FP), DI
	feMov(AX,BX,CX,DX, 0(DI),8(DI),16(DI),24(DI))
	RET

// func feSub(c, a, b *gfP2)
TEXT ·feSub(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	feMov(0(SI),8(SI),16(SI),24(SI), AX,BX,CX,DX)
	feSub(0(DI),8(DI),16(DI),24(DI), AX,BX,CX,DX)

	MOVQ c+0(FP), DI
	feMov(AX,BX,CX,DX, 0(DI),8(DI),16(DI),24(DI))
	RET

// func feMul(c, a, b *fieldElem)
TEXT ·feMul(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI

	feMul(CX, 0(DI),8(DI),16(DI),24(DI), 0(SI),8(SI),16(SI),24(SI), R8,R9,R10,R11)

	MOVQ c+0(FP), DI
	MOVQ  R8,  0(DI)
	MOVQ  R9,  8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
	RET

// func feSquare(c, a *fieldElem)
TEXT ·feSquare(SB),0,$0-16
	MOVQ a+8(FP), DI

	feSquare(CX,R12,R13, 0(DI),8(DI),16(DI),24(DI), R8,R9,R10,R11)

	MOVQ c+0(FP), DI
	MOVQ R8, 0(DI)
	MOVQ R9, 8(DI)
	MOVQ R10, 16(DI)
	MOVQ R11, 24(DI)
	RET

// func feInvert(c, a *gfP2)
TEXT ·feInvert(SB),0,$0-16
	MOVQ a+8(FP), DI

	bfeSquare(CX, 0(DI),8(DI), R12,R13)
	bfeSquareAdd(CX, 16(DI),24(DI), R12,R13)
	bfeMulReduce(CX, R12,R13)

	// Invert the value in y=R12:R13. Store output in a=R8:R9.
	//
	// This is a GCD-based algorithm for inversion from "Prime Numbers: A
	// Computational Perspective" by Crandall, Pomerance. Section 9.4.2.
	MOVQ $1, R8
	MOVQ $0, R9 // a
	MOVQ $0, R10
	MOVQ $0, R11 // b
	MOVQ $0, R14
	MOVQ $0, R15 // z
	bfeNeg(R14,R15) // Set z to p=2^127-1.

reduce:
	// Find e such that 2^e||y.
	BSFQ R12, CX
	JNZ common
	BSFQ R13, CX // Handle shift greater than 64-bits.
	JZ end

	// Shift y right 64 bits.
	MOVQ R13, R12
	MOVQ $0, R13

	// Rotate a right 64 bits.
	SHLQ $1, R9
	XCHGQ R8, R9
	SHRQ $1, R8:R9
	SHRQ $1, R9

common:
	// Shift trailing zeroes off of y.
	SHRQ CL, R12:R13
	SHRQ CL, R13

	// a = 2^(127-e) * a (Right rotation).
	MOVQ $0, BX
	SHRQ CL, BX:R8
	SHRQ CL, R8:R9
	SHRQ CL, R9

	SHRQ $1, BX
	XORQ BX, R9

	// if (y == 1) return a
	CMPQ R12, $1
	JNE cont
	CMPQ R13, $0
	JE end

cont:
	// (a, b, y, z) = (a+b, a, y+z, y)
	MOVQ R9, AX
	MOVQ R13, BX

	XADDQ R10, R8
	ADCQ R11, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9

	XADDQ R14, R12
	ADCQ R15, R13

	MOVQ AX, R11
	MOVQ BX, R15

	JMP reduce

end:
	MOVQ c+0(FP), SI

	bfeMul(R14, R8,R9, 0(DI),8(DI), R12,R13)
	bfeMulReduce(R14, R12,R13)
	MOVQ R12, 0(SI)
	MOVQ R13, 8(SI)

	bfeNeg(R8,R9)
	bfeMul(R14, R8,R9, 16(DI),24(DI), R12,R13)
	bfeMulReduce(R14, R12,R13)
	MOVQ R12, 16(SI)
	MOVQ R13, 24(SI)

	RET
