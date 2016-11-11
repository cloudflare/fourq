// func feMul(c, a, b *gfP2)
TEXT ·feMul(SB),0,$0-24
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI

// Mult a.x * b.x = 0(DI)||8(DI) * 0(SI)||8(SI). Store in (R9*2^64 + R8).
	MOVQ $0, R15 // R15 is going to be used for super-carries.

	// Compute lower block of multiplication. Store in (R9*2^64 + R8).
	MOVQ 0(DI), AX
	MULQ 0(SI)
	MOVQ AX, R8
	MOVQ DX, BX

	MOVQ 0(DI), AX
	MULQ 8(SI)
	MOVQ DX, CX
	ADDQ AX, BX
	ADCQ $0, CX
	ADCQ $0, R15

	MOVQ 8(DI), AX
	MULQ 0(SI)
	ADDQ AX, BX
	ADCQ DX, CX
	ADCQ $0, R15
	MOVQ BX, R9

	// Compute upper block of multiplication. Store in (DX*2^64 + AX).
	MOVQ 8(DI), AX
	MULQ 8(SI)
	ADDQ CX, AX
	ADCQ R15, DX // 8(DI) and 8(SI) are 62-bit, so top two bits of DX are zero. (Only one now.)

	// Shift upper block by 1.
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX // Goes into zero at bottom of DX.

	// Add upper and lower blocks together.
	MOVQ $0, R15 // Clear R15 -- next ADCQ below might cause super-carry.

	ADDQ AX, R8
	ADCQ DX, R9
	ADCQ $0, R15

	SHLQ $1, R15
	BTRQ $63, R9
	ADCQ $0, R15

	ADDQ R15, R8
	ADCQ $0, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9

// Mult a.y * b.y = 16(DI)||24(DI) * 16(SI)||24(SI). Leave in (DX*2^64 + AX).
// (Code compressed.)
	MOVQ $0, R15
	MOVQ 16(DI), AX
	MULQ 16(SI)
	MOVQ AX, R12
	MOVQ DX, BX
	MOVQ 16(DI), AX
	MULQ 24(SI)
	MOVQ DX, CX
	ADDQ AX, BX
	ADCQ $0, CX
	ADCQ $0, R15
	MOVQ 24(DI), AX
	MULQ 16(SI)
	ADDQ AX, BX
	ADCQ DX, CX
	ADCQ $0, R15
	MOVQ BX, R13
	MOVQ 24(DI), AX
	MULQ 24(SI)
	ADDQ CX, AX
	ADCQ R15, DX
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX
	MOVQ $0, R15
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

// Sub (DX*2^64 + AX) from (R8*2^64 + R9).
	NOTQ AX
	NOTQ DX
	BTRQ $63, DX

	ADDQ AX, R8
	ADCQ DX, R9

	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9

// Mult a.x * b.y = 0(DI)||8(DI) * 16(SI)||24(SI). Store in (R11*2^64 + R10).
// (Code compressed.)
	MOVQ $0, R15
	MOVQ 0(DI), AX
	MULQ 16(SI)
	MOVQ AX, R10
	MOVQ DX, BX
	MOVQ 0(DI), AX
	MULQ 24(SI)
	MOVQ DX, CX
	ADDQ AX, BX
	ADCQ $0, CX
	ADCQ $0, R15
	MOVQ 8(DI), AX
	MULQ 16(SI)
	ADDQ AX, BX
	ADCQ DX, CX
	ADCQ $0, R15
	MOVQ BX, R11
	MOVQ 8(DI), AX
	MULQ 24(SI)
	ADDQ CX, AX
	ADCQ R15, DX
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX
	MOVQ $0, R15
	ADDQ AX, R10
	ADCQ DX, R11
	ADCQ $0, R15
	SHLQ $1, R15
	BTRQ $63, R11
	ADCQ $0, R15
	ADDQ R15, R10
	ADCQ $0, R11
	BTRQ $63, R11
	ADCQ $0, R10
	ADCQ $0, R11

// Mult a.y * b.x = 16(DI)||24(DI) * 0(SI)||8(SI). Leave in (DX*2^64 + AX).
// (Code compressed.)
	MOVQ $0, R15
	MOVQ 16(DI), AX
	MULQ 0(SI)
	MOVQ AX, R12
	MOVQ DX, BX
	MOVQ 16(DI), AX
	MULQ 8(SI)
	MOVQ DX, CX
	ADDQ AX, BX
	ADCQ $0, CX
	ADCQ $0, R15
	MOVQ 24(DI), AX
	MULQ 0(SI)
	ADDQ AX, BX
	ADCQ DX, CX
	ADCQ $0, R15
	MOVQ BX, R13
	MOVQ 24(DI), AX
	MULQ 8(SI)
	ADDQ CX, AX
	ADCQ R15, DX
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX
	MOVQ $0, R15
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

// Add (R11*2^64 + R10) to (DX*2^64 + AX).
	ADDQ R10, AX
	ADCQ R11, DX

	BTRQ $63, DX
	ADCQ $0, AX
	ADCQ $0, DX

// Move out and return.
	MOVQ c+0(FP), DI

	MOVQ R8,  0(DI)
	MOVQ R9,  8(DI)
	MOVQ AX, 16(DI)
	MOVQ DX, 24(DI)

	RET
