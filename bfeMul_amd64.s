// +build amd64,!noasm

// func bfeMul(c, a, b *baseFieldElem)
TEXT ·bfeMul(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI

	MOVQ 0(DI), R8
	MOVQ 8(DI), R9
	MOVQ 0(SI), R10
	MOVQ 8(SI), R11

	MOVQ $0, R15 // R15 is going to be used for super-carries.

	// Compute lower block of multiplication. Store in (R13*2^64 + R12).
	MOVQ R8, AX
	MULQ R10
	MOVQ AX, R12
	MOVQ DX, BX

	MOVQ R8, AX
	MULQ R11
	MOVQ DX, CX
	ADDQ AX, BX
	ADCQ $0, CX
	ADCQ $0, R15

	MOVQ R9, AX
	MULQ R10
	ADDQ AX, BX
	ADCQ DX, CX
	ADCQ $0, R15
	MOVQ BX, R13

	// Compute upper block of multiplication. Store in (DX*2^64 + AX).
	MOVQ R9, AX
	MULQ R11
	ADDQ CX, AX
	ADCQ R15, DX // R9 and R11 are 62-bit, so top two bits of DX are zero. (Only one now.)

	// Shift upper block by 1.
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX // Goes into zero at bottom of DX.

	// Add upper and lower blocks together.
	MOVQ $0, R15 // Clear R15 -- next ADCQ below might cause super-carry.

	ADDQ R12, AX
	ADCQ R13, DX
	ADCQ $0, R15

	SHLQ $1, R15
	BTRQ $63, DX
	ADCQ $0, R15

	ADDQ R15, AX
	ADCQ $0, DX
	BTRQ $63, DX
	ADCQ $0, AX
	ADCQ $0, DX

	// TODO(brendan): Final reduction.

	// Store output.
	MOVQ c+0(FP), DI
	MOVQ AX, 0(DI)
	MOVQ DX, 8(DI)

	RET
