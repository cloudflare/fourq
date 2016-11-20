#include "field.h"

// func feDbl(c, a *fieldElem)
TEXT ·feDbl(SB),0,$0-16
	MOVQ a+8(FP), DI
	feMov(0(DI),8(DI),16(DI),24(DI), AX,BX,CX,DX)
	feDbl(AX,BX,CX,DX)

	MOVQ c+0(FP), DI
	feMov(AX,BX,CX,DX, 0(DI),8(DI),16(DI),24(DI))
	RET

// func feAdd(c, a, b *fieldElem)
TEXT ·feAdd(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	feMov(0(DI),8(DI),16(DI),24(DI), AX,BX,CX,DX)
	feAdd(0(SI),8(SI),16(SI),24(SI), AX,BX,CX,DX)

	MOVQ c+0(FP), DI
	feMov(AX,BX,CX,DX, 0(DI),8(DI),16(DI),24(DI))
	RET

// func feSub(c, a, b *fieldElem)
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
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))
	RET

// func feSquare(c, a *fieldElem)
TEXT ·feSquare(SB),0,$0-16
	MOVQ a+8(FP), DI
	feSquare(CX,R12,R13, 0(DI),8(DI),16(DI),24(DI), R8,R9,R10,R11)

	MOVQ c+0(FP), DI
	feMov(R8,R9,R10,R11, 0(DI),8(DI),16(DI),24(DI))
	RET
