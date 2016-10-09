// +build amd64,!noasm

// func bfeAdd(dst, a, b *baseFieldElem)
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
