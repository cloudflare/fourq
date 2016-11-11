// func bfeAdd(c, a, b *baseFieldElem)
TEXT ·bfeDbl(SB),0,$0-16
	MOVQ a+8(FP), DI

	MOVQ 0(DI), AX
	MOVQ 8(DI), BX

	SHLQ $1, BX
	SHLQ $1, AX
	ADCQ $0, BX

	BTRQ $63, BX
	ADCQ $0, AX
	ADCQ $0, BX

	MOVQ c+0(FP), DI
	MOVQ AX, 0(DI)
	MOVQ BX, 8(DI)
	RET

// func bfeAdd(c, a, b *baseFieldElem)
TEXT ·bfeAdd(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI

	MOVQ 0(DI), AX
	MOVQ 8(DI), BX
	ADDQ 0(SI), AX
	ADCQ 8(SI), BX

	BTRQ $63, BX
	ADCQ $0, AX
	ADCQ $0, BX

	MOVQ c+0(FP), DI
	MOVQ AX, 0(DI)
	MOVQ BX, 8(DI)
	RET

// func bfeSub(c, a, b *baseFieldElem)
TEXT ·bfeSub(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI

	MOVQ 0(SI), AX
	MOVQ 8(SI), BX
	NOTQ AX
	NOTQ BX
	BTRQ $63, BX

	ADDQ 0(DI), AX
	ADCQ 8(DI), BX

	BTRQ $63, BX
	ADCQ $0, AX
	ADCQ $0, BX

	MOVQ c+0(FP), DI
	MOVQ AX, 0(DI)
	MOVQ BX, 8(DI)
	RET
