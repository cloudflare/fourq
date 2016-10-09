
// func bfeSquare(dst, a *baseFieldElem)
TEXT ·bfeSquare(SB), $0-16 // TODO(brendan): Frame size annotations.
	MOVQ a+8(FP), SI

	// TODO(brendan): Should I save the contents of these registers and restore them at the end?
	MOVQ 0(SI), R8
	MOVQ 8(SI), R9

	MOVQ $0, R15 // R15 is going to be used for super-carries.

	// Compute lower block of multiplication. Store in (R11*2^64 + R10).
	MOVQ R8, AX
	MULQ AX
	MOVQ AX, R10
	MOVQ DX, CX

	MOVQ R8, AX
	MULQ R9
	SHLQ $1, DX // Top bit of DX was zero, because top bit of R9 is. Now, not necessarily.
	SHLQ $1, AX
	ADCQ $0, DX // This goes into the zero at the bottom of DX.

	ADDQ CX, AX // Add carry from previous mult.
	ADCQ $0, DX
	ADCQ $0, R15

	MOVQ AX, R11
	MOVQ DX, CX

	// Compute upper block of multiplication. Store in (DX*2^64 + AX).
	MOVQ R9, AX
	MULQ AX
	ADDQ CX, AX
	ADCQ R15, DX // R9 is 62-bit, so top two bits of DX are zero. (Only one now.)

	// Shift upper block by 1.
	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX // Goes into zero at bottom of DX.

	// Add upper and lower blocks together.
	MOVQ $0, R15 // Clear R15 -- next ADCQ below might cause super-carry.

	ADDQ R10, AX
	ADCQ R11, DX
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
	MOVQ c+0(FP), SI
	MOVQ AX, 0(SI)
	MOVQ DX, 8(SI)

	RET
