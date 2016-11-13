// func feInvert(c, a *gfP2)
TEXT ·feInvert(SB),0,$0-16
	MOVQ a+8(FP), DI

// bfeSquare(t, &a.x)
	MOVQ 0(DI), R8
	MOVQ 8(DI), R9

	MOVQ $0, R14 // R14 is going to be used for super-carries.

	// One
	MOVQ R8, AX
	MULQ R8
	MOVQ AX, R12
	MOVQ DX, R13

	// Two
	MOVQ R8, AX
	MULQ R9
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Four
	MOVQ R9, AX
	MULQ R9

	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX

	ADDQ AX, R12
	ADCQ DX, R13
	ADCQ $0, R14

	// Reduce
	SHLQ $1, R14
	BTRQ $63, R13
	ADCQ R14, R12
	ADCQ $0, R13

	BTRQ $63, R13
	ADCQ $0, R12
	ADCQ $0, R13

// bfeSquare(t, &a.y)
// bfeAdd(t, t, t2)
	MOVQ 16(DI), R8
	MOVQ 24(DI), R9

	MOVQ $0, R14 // R14 is going to be used for super-carries.

	// One
	MOVQ R8, AX
	MULQ R8
	ADDQ AX, R12
	ADCQ DX, R13
	ADCQ $0, R14

	// Two
	MOVQ R8, AX
	MULQ R9
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Four
	MOVQ R9, AX
	MULQ R9

	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX

	ADDQ AX, R12
	ADCQ DX, R13
	ADCQ $0, R14

	// Reduce
	SHLQ $1, R14
	BTRQ $63, R13
	ADCQ R14, R12
	ADCQ $0, R13

	BTRQ $63, R13
	ADCQ $0, R12
	ADCQ $0, R13

// bfeInvert(inv, t)
	MOVQ $1, R8
	MOVQ $0, R9 // a
	MOVQ $0, R10
	MOVQ $0, R11 // b
	// MOVQ 0(DI), R12
	// MOVQ 8(DI), R13 // y
	MOVQ $0, R14
	MOVQ $0, R15 // z

	NOTQ R14
	NOTQ R15
	BTRQ $63, R15

reduce:
	// Find e such that 2^e||y.
	BSFQ R12, CX
	JNZ common
	BSFQ R13, CX // Handle shift greater than 64-bits.
	JZ end

	// Shift y right 64 bits.
	MOVQ R13, R12
	MOVQ $0, R13

	// Rotate a right 64 bits.
	SHLQ $1, R9
	XCHGQ R8, R9
	SHRQ $1, R8:R9
	SHRQ $1, R9

common:
	// Shift trailing zeroes off of y.
	SHRQ CL, R12:R13
	SHRQ CL, R13

	// a = 2^(127-e) * a (Right rotation).
	MOVQ $0, BX
	SHRQ CL, BX:R8
	SHRQ CL, R8:R9
	SHRQ CL, R9

	SHRQ $1, BX
	XORQ BX, R9

	// if (y == 1) return a
	CMPQ R12, $1
	JNE cont
	CMPQ R13, $0
	JE end

cont:
	// (a, b, y, z) = (a+b, a, y+z, y)
	MOVQ R9, AX
	MOVQ R13, BX

	XADDQ R10, R8
	ADCQ R11, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9

	XADDQ R14, R12
	ADCQ R15, R13

	MOVQ AX, R11
	MOVQ BX, R15

	JMP reduce

end:
	// Return a.
	// MOVQ c+0(FP), DI
	// MOVQ R8, 0(DI)
	// MOVQ R9, 8(DI)
	MOVQ c+0(FP), SI

// bfeMul(&e.x, &a.x, inv)
	MOVQ 0(DI), R10
	MOVQ 8(DI), R11

	MOVQ $0, R14 // R14 is going to be used for super-carries.

	// One
	MOVQ R8, AX
	MULQ R10
	MOVQ AX, R12
	MOVQ DX, R13

	// Two
	MOVQ R8, AX
	MULQ R11
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Three
	MOVQ R9, AX
	MULQ R10
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Four
	MOVQ R9, AX
	MULQ R11

	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX

	ADDQ AX, R12
	ADCQ DX, R13
	ADCQ $0, R14

	// Reduce
	SHLQ $1, R14
	BTRQ $63, R13
	ADCQ R14, R12
	ADCQ $0, R13

	BTRQ $63, R13
	ADCQ $0, R12
	ADCQ $0, R13

	// Move out.
	MOVQ R12, 0(SI)
	MOVQ R13, 8(SI)

// e.y.Neg(&a.y)
// bfeMul(&e.y, &e.y, inv)
	MOVQ 16(DI), R10
	MOVQ 24(DI), R11

	NOTQ R10
	NOTQ R11
	BTRQ $63, R11

	MOVQ $0, R14 // R14 is going to be used for super-carries.

	// One
	MOVQ R8, AX
	MULQ R10
	MOVQ AX, R12
	MOVQ DX, R13

	// Two
	MOVQ R8, AX
	MULQ R11
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Three
	MOVQ R9, AX
	MULQ R10
	SHLQ $1, DX
	ADDQ DX, R12
	ADCQ AX, R13
	ADCQ $0, R14

	// Four
	MOVQ R9, AX
	MULQ R11

	SHLQ $1, DX
	SHLQ $1, AX
	ADCQ $0, DX

	ADDQ AX, R12
	ADCQ DX, R13
	ADCQ $0, R14

	// Reduce
	SHLQ $1, R14
	BTRQ $63, R13
	ADCQ R14, R12
	ADCQ $0, R13

	BTRQ $63, R13
	ADCQ $0, R12
	ADCQ $0, R13

	// Move out.
	MOVQ R12, 16(SI)
	MOVQ R13, 24(SI)

	RET
