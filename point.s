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

	// Layout of Stack: D || G || F

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
