#include "field.h"

// func pDbl(a *point)
TEXT ·pDbl(SB),0,$128-8
	MOVQ a+0(FP), DI

	feSquare(CX,R12,R13, 0(DI),8(DI),16(DI),24(DI), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(SP),8(SP),16(SP),24(SP))

	feSquare(CX,R12,R13, 32(DI),40(DI),48(DI),56(DI), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	feSquare(CX,R12,R13, 96(DI),104(DI),112(DI),120(DI), R8,R9,R10,R11)
	feDbl(R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 64(SP),72(SP),80(SP),88(SP))

	// E := newGFp2().Add(&a.x, &a.y)
	feMov(0(DI),8(DI),16(DI),24(DI), BX,SI,R14,R15)
	feAdd(32(DI),40(DI),48(DI),56(DI), BX,SI,R14,R15)

	feSquare(CX,R12,R13, BX,SI,R14,R15, R8,R9,R10,R11)

	// E.Sub(E, A).Sub(E, B)
	feMov(0(SP),8(SP),16(SP),24(SP), AX,BX,CX,DX)
	feReverseSub(R8,R9,R10,R11, AX,BX,CX,DX)

	feMov(32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15)
	feReverseSub(R8,R9,R10,R11, R12,R13,R14,R15)

	// Current layout of memory:
	// AX, BX, CX, DX:     -A
	// R8, R9, R10, R11:    E
	// R12, R13, R14, R15: -B

	feMov(R8,R9,R10,R11, 96(SP),104(SP),112(SP),120(SP))

	// Current layout of stack:
	// A, B, C, E

	// Load B into R8...R11.
	feMov(32(SP),40(SP),48(SP),56(SP), R8,R9,R10,R11)

	// Calculate G := newGFp2().Sub(B, A) into R8...R11.
	feAdd(AX,BX,CX,DX, R8,R9,R10,R11)

	// H := newGFp2().Add(B, A)
	// H.Neg(H)
	// Calculate H into AX...DX.
	feAdd(R12,R13,R14,R15, AX,BX,CX,DX)

	feMov(AX,BX, CX, DX,  0(SP), 8(SP),16(SP),24(SP))
	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	// Load C into R12...R15
	feMov(64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15)

	// Calculate F := newGFp2().Sub(G, C) into R12...R15
	feSub(R8,R9,R10,R11, R12,R13,R14,R15)

	// Current layout of memory:
	// AX, BX, CX, DX:     H
	// R8, R9, R10, R11:   G
	// R12, R13, R14, R15: F
	//
	// Current layout of stack:
	// H, G, C, E

	feMul(CX, 96(SP),104(SP),112(SP),120(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))

	feMul(CX, 32(SP),40(SP),48(SP),56(SP), 0(SP),8(SP),16(SP),24(SP), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(DI),40(DI),48(DI),56(DI))

	feMul(CX, 96(SP),104(SP),112(SP),120(SP), 0(SP),8(SP),16(SP),24(SP), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 64(DI),72(DI),80(DI),88(DI))

	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 96(DI),104(DI),112(DI),120(DI))

	RET
