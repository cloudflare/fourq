#include "field.h"

// func pDbl(a *point)
TEXT ·pDbl(SB),0,$96-8
	MOVQ a+0(FP), DI

	// A = X1^2
	feSquare(CX,BX,SI, 0(DI),8(DI),16(DI),24(DI), R8,R9,R10,R11)

	// B = Y1^2
	feSquare(CX,BX,SI, 32(DI),40(DI),48(DI),56(DI), R12,R13,R14,R15)

	// D = -(A + B)
	// G = B - A
	feMov(R12,R13,R14,R15, AX,BX, CX, DX)
	feNeg( AX, BX, CX, DX)
	feNeg( R8, R9,R10,R11)
	feAdd( R8, R9,R10,R11, AX,BX, CX, DX) // D
	feAdd(R12,R13,R14,R15, R8,R9,R10,R11) // G

	feMov(AX,BX,CX,DX, 0(SP),8(SP),16(SP),24(SP))

	// F = G - 2*Z1^2
	feSquare(CX,BX,SI, 64(DI),72(DI),80(DI),88(DI), R12,R13,R14,R15)
	feDbl(R12,R13,R14,R15)
	feSub(R8,R9,R10,R11, R12,R13,R14,R15)

	feMov( R8, R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))
	feMov(R12,R13,R14,R15, 64(SP),72(SP),80(SP),88(SP))

	// E = (X1 + Y1)^2 + D
	feMov( 0(DI), 8(DI),16(DI),24(DI), R8,R9,R10,R11)
	feAdd(32(DI),40(DI),48(DI),56(DI), R8,R9,R10,R11)

	feSquare(CX,BX,SI, R8,R9,R10,R11, R12,R13,R14,R15)
	feAdd(0(SP),8(SP),16(SP),24(SP), R12,R13,R14,R15)

	// Layout of stack: D || G || F

	// X3 = F * E
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))

	// T3 = D * E
	feMul(CX, 0(SP),8(SP),16(SP),24(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 96(DI),104(DI),112(DI),120(DI))

	feMov(32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15)

	// Y3 = D * G
	feMul(CX, 0(SP),8(SP),16(SP),24(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(DI),40(DI),48(DI),56(DI))

	// Z3 = F * G
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 64(DI),72(DI),80(DI),88(DI))

	RET

// func pMixedAdd(a, b *point)
TEXT ·pMixedAdd(SB),0,$96-16
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI

	// A = X1 * X2
	feMul(CX, 0(DI),8(DI),16(DI),24(DI), 0(SI),8(SI),16(SI),24(SI), R8,R9,R10,R11)

	// B = Y1 * Y2
	feMul(CX, 32(DI),40(DI),48(DI),56(DI), 32(SI),40(SI),48(SI),56(SI), R12,R13,R14,R15)

	// D = A + B
	feAdd(R8,R9,R10,R11, R12,R13,R14,R15)
	feMov(R12,R13,R14,R15, 0(SP),8(SP),16(SP),24(SP))

	// E = (X1 + Y1)(X2 + Y2) - D
	feMov( 0(DI), 8(DI),16(DI),24(DI), R8,R9,R10,R11)
	feAdd(32(DI),40(DI),48(DI),56(DI), R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	feMov( 0(SI), 8(SI),16(SI),24(SI), R8,R9,R10,R11)
	feAdd(32(SI),40(SI),48(SI),56(SI), R8,R9,R10,R11)

	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R8,R9,R10,R11, R12,R13,R14,R15)
	feMov(0(SP),8(SP),16(SP),24(SP), R8,R9,R10,R11)
	feSub(R12,R13,R14,R15, R8,R9,R10,R11)

	feMov(R8,R9,R10,R11, 32(SP),40(SP),48(SP),56(SP))

	// C = T1 * T2
	feMul(CX, 96(DI),104(DI),112(DI),120(DI), 96(SI),104(SI),112(SI),120(SI), R12,R13,R14,R15)

	// G = Z1 + C
	// F = Z1 - C
	feMov(64(DI),72(DI),80(DI),88(DI), AX,BX,CX,DX)
	feMov(AX,BX,CX,DX, R8,R9,R10,R11)

	feAdd(R12,R13,R14,R15,  AX, BX, CX, DX) // G
	feSub( R8, R9,R10,R11, R12,R13,R14,R15) // F

	feMov(AX,BX,CX,DX, 64(SP),72(SP),80(SP),88(SP))

	// Layout of stack: D || E || G

	// X3 = E * F
	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))

	// Z3 = G * F
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 64(DI),72(DI),80(DI),88(DI))

	feMov(0(SP),8(SP),16(SP),24(SP), R12,R13,R14,R15)

	// Y3 = G * D
	feMul(CX, 64(SP),72(SP),80(SP),88(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 32(DI),40(DI),48(DI),56(DI))

	// T3 = E * D
	feMul(CX, 32(SP),40(SP),48(SP),56(SP), R12,R13,R14,R15, R8,R9,R10,R11)
	feMov(R8,R9,R10,R11, 96(DI),104(DI),112(DI),120(DI))

	RET
