#include "macros.s"

#define x 0(DI),8(DI),16(DI),24(DI)

// func pMixedAdd(a, b *point)
TEXT ·pMixedAdd(SB),0,$96-16
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI

	// A := newGFp2().Sub(&a.y, &a.x)
	feMov( 0(DI), 8(DI),16(DI),24(DI), AX,BX,CX,DX)
	feSub(32(DI),40(DI),48(DI),56(DI), AX,BX,CX,DX)
	feMov(AX,BX,CX,DX, 0(SP),8(SP),16(SP),24(SP))

	// tmp := newGFp2().Add(&b.y, &b.x)
	feMov( 0(SI), 8(SI),16(SI),24(SI), R12,R13,R14,R15)
	feAdd(32(SI),40(SI),48(SI),56(SI), R12,R13,R14,R15)

	// feMul(A, A, tmp)
	feMul(CX, 0(SP),8(SP),16(SP),24(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(SP),8(SP),16(SP),24(SP))

	// B := newGFp2().Sub(&b.y, &b.x)
	feMov( 0(SI), 8(SI),16(SI),24(SI), AX,BX,CX,DX)
	feSub(32(SI),40(SI),48(SI),56(SI), AX,BX,CX,DX)
	feMov(AX,BX,CX,DX, 32(SP),40(SP),48(SP),56(SP))

	// tmp.Add(&a.y, &a.x)
	feMov( 0(DI), 8(DI),16(DI),24(DI), R12,R13,R14,R15)
	feAdd(32(DI),40(DI),48(DI),56(DI), R12,R13,R14,R15)

	// feMul(B, B, tmp)
	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	// C := newGFp2()
	// feMul(C, &a.z, &b.t)
	// C.Dbl(C)
	feMul(CX, 96(DI),104(DI),112(DI),120(DI), 64(SI),72(SI),80(SI),88(SI), R8,R9,R10,R11)
	feDbl(R8,R9,R10,R11)

	// D := newGFp2().Dbl(&a.t)
	feMov(64(DI),72(DI),80(DI),88(DI), R12,R13,R14,R15)
	feDbl(R12,R13,R14,R15)

	// Current register layout:
	// R8, R9, R10, R11: C
	// R12, R13, R14, R15: D
	//
	// Current stack layout:
	// A, B

	// E := newGFp2().Add(D, C)
	feMov(R8,R9,R10,R11, AX,BX,CX,DX)
	feAdd(R12,R13,R14,R15, AX,BX,CX,DX)
	feMov(AX,BX,CX,DX, 64(SP),72(SP),80(SP),88(SP))

	// H := newGFp2().Sub(D, C)
	feReverseSub(R12,R13,R14,R15, R8,R9,R10,R11)

	// F := newGFp2().Sub(B, A)
	// G := newGFp2().Add(B, A)
	feMov(0(SP),8(SP),16(SP),24(SP), AX,BX,CX,DX)
	feMov(32(SP),40(SP),48(SP),56(SP), R8,R9,R10,R11)

	feAdd(AX,BX,CX,DX, R8,R9,R10,R11)
	feSub(32(SP),40(SP),48(SP),56(SP), AX,BX,CX,DX)

	feMov(AX,BX,CX,DX, 0(SP),8(SP),16(SP),24(SP))
	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	// Current register layout:
	// R12, R13, R14, R15: H
	//
	// Current stack layout:
	// F, G, E

	// feMul(&c.x, E, F)
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), 0(SP),8(SP),16(SP),24(SP), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))

	// feMul(&c.y, G, H)
	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(DI),40(DI),48(DI),56(DI))

	// feMul(&c.t, E, H)
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 64(DI),72(DI),80(DI),88(DI))

	// feMul(&c.z, F, G)
	feMul(CX, 0(SP),8(SP),16(SP),24(SP), 32(SP),40(SP),48(SP),56(SP), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 96(DI),104(DI),112(DI),120(DI))

	RET
