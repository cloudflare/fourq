#include "base.h"

// func bfeAdd(c, a, b *baseFieldElem)
TEXT ·bfeAdd(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	bfeMov(0(DI),8(DI), AX,BX)
	bfeAdd(0(SI),8(SI), AX,BX)

	MOVQ c+0(FP), DI
	bfeMov(AX,BX, 0(DI),8(DI))
	RET

// func bfeSub(c, a, b *baseFieldElem)
TEXT ·bfeSub(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	bfeMov(0(SI),8(SI), AX,BX)
	bfeSub(0(DI),8(DI), AX,BX)

	MOVQ c+0(FP), DI
	bfeMov(AX,BX, 0(DI),8(DI))
	RET

// func bfeMul(c, a, b *baseFieldElem)
TEXT ·bfeMul(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	bfeMul(CX, 0(DI),8(DI), 0(SI),8(SI), R8,R9)

	MOVQ c+0(FP), DI
	bfeMov(R8,R9, 0(DI),8(DI))
	RET

// func bfeSquare(c, a *baseFieldElem)
TEXT ·bfeSquare(SB),0,$0-16
	MOVQ a+8(FP), DI
	bfeSquare(CX, 0(DI),8(DI), R8,R9)

	MOVQ c+0(FP), DI
	bfeMov(R8,R9, 0(DI),8(DI))
	RET
