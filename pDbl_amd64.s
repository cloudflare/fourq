// One base-field element is 16 bytes
// One field element is two base-field elements = 32 bytes
// One point is two field elements              = 64 bytes

// func pDbl(a *point)
TEXT ·pDbl(SB),0,$128-8
	MOVQ a+0(FP), DI

	// feSquare(A, &a.x)
		// Compute a.x + a.y = 0(DI)||8(DI) + 16(DI)||24(DI). Store in (R11*2^64+R10).
		MOVQ 0(DI), R10
		MOVQ 8(DI), R11
		ADDQ 16(DI), R10
		ADCQ 24(DI), R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Compute a.x - a.y = 0(DI)||8(DI) - 16(DI)||24(DI). Store in (R13*2^64+R12).
		MOVQ 16(DI), R12
		MOVQ 24(DI), R13

		NOTQ R12
		NOTQ R13
		BTRQ $63, R13

		ADDQ 0(DI), R12
		ADCQ 8(DI), R13
		BTRQ $63, R13
		ADCQ $0, R12
		ADCQ $0, R13

		// Mult (R11*2^64+R10) * (R13*2^64+R12). Store in (R9*2^64+R8).
		MOVQ $0, CX
		MOVQ R10, AX
		MULQ R12
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ R10, AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * a.y = 0(DI)||8(DI) * 16(DI)||24(DI). Store in (R11*2^64+R10).
		MOVQ $0, CX
		MOVQ 0(DI), AX
		MULQ 16(DI)
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 0(DI), AX
		MULQ 24(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 8(DI), AX
		MULQ 16(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 8(DI), AX
		MULQ 24(DI)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Double (R11*2^64+R10) in-place.
		SHLQ $1, R11
		SHLQ $1, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10

		// Move out.
		MOVQ R8, 0(SP)
		MOVQ R9, 8(SP)
		MOVQ R10, 16(SP)
		MOVQ R11, 24(SP)

	// feSquare(B, &a.y)
		// Compute a.x + a.y = 32(DI)||40(DI) + 48(DI)||56(DI). Store in (R11*2^64+R10).
		MOVQ 32(DI), R10
		MOVQ 40(DI), R11
		ADDQ 48(DI), R10
		ADCQ 56(DI), R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Compute a.x - a.y = 32(DI)||40(DI) - 48(DI)||56(DI). Store in (R13*2^64+R12).
		MOVQ 48(DI), R12
		MOVQ 56(DI), R13

		NOTQ R12
		NOTQ R13
		BTRQ $63, R13

		ADDQ 32(DI), R12
		ADCQ 40(DI), R13
		BTRQ $63, R13
		ADCQ $0, R12
		ADCQ $0, R13

		// Mult (R11*2^64+R10) * (R13*2^64+R12). Store in (R9*2^64+R8).
		MOVQ $0, CX
		MOVQ R10, AX
		MULQ R12
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ R10, AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * a.y = 32(DI)||40(DI) * 48(DI)||56(DI). Store in (R11*2^64+R10).
		MOVQ $0, CX
		MOVQ 32(DI), AX
		MULQ 48(DI)
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 32(DI), AX
		MULQ 56(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(DI), AX
		MULQ 48(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(DI), AX
		MULQ 56(DI)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Double (R11*2^64+R10) in-place.
		SHLQ $1, R11
		SHLQ $1, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10

		// Move out.
		MOVQ R8, 32(SP)
		MOVQ R9, 40(SP)
		MOVQ R10, 48(SP)
		MOVQ R11, 56(SP)

	// feSquare(C, &a.z)
	// C.Dbl(C)
		// Compute a.x + a.y = 96(DI)||104(DI) + 112(DI)||120(DI). Store in (R11*2^64+R10).
		MOVQ 96(DI), R10
		MOVQ 104(DI), R11
		ADDQ 112(DI), R10
		ADCQ 120(DI), R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Compute a.x - a.y = 96(DI)||104(DI) - 112(DI)||120(DI). Store in (R13*2^64+R12).
		MOVQ 112(DI), R12
		MOVQ 120(DI), R13

		NOTQ R12
		NOTQ R13
		BTRQ $63, R13

		ADDQ 96(DI), R12
		ADCQ 104(DI), R13
		BTRQ $63, R13
		ADCQ $0, R12
		ADCQ $0, R13

		// Mult (R11*2^64+R10) * (R13*2^64+R12). Store in (R9*2^64+R8).
		MOVQ $0, CX
		MOVQ R10, AX
		MULQ R12
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ R10, AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * a.y = 96(DI)||104(DI) * 112(DI)||120(DI). Store in (R11*2^64+R10).
		MOVQ $0, CX
		MOVQ 96(DI), AX
		MULQ 112(DI)
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 96(DI), AX
		MULQ 120(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(DI), AX
		MULQ 112(DI)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(DI), AX
		MULQ 120(DI)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Double (R11*2^64+R10) in-place.
		SHLQ $1, R11
		SHLQ $1, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Double entire element.
		SHLQ $1, R9
		SHLQ $1, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9
		SHLQ $1, R11
		SHLQ $1, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Move out.
		MOVQ R8, 64(SP)
		MOVQ R9, 72(SP)
		MOVQ R10, 80(SP)
		MOVQ R11, 88(SP)

	// E := newGFp2().Add(&a.x, &a.y)
	MOVQ 0(DI), BX
	MOVQ 8(DI), SI
	MOVQ 16(DI), R14
 	MOVQ 24(DI), R15

	ADDQ 32(DI), BX
	ADCQ 40(DI), SI
	BTRQ $63, SI
	ADCQ $0, BX
	ADCQ $0, SI

	ADDQ 48(DI), R14
	ADCQ 56(DI), R15
	BTRQ $63, R15
	ADCQ $0, R14
	ADCQ $0, R15

	// feSquare(E, E)
		// Compute a.x + a.y = BX||SI + R14||R15. Store in (R11*2^64+R10).
		MOVQ BX, R10
		MOVQ SI, R11
		ADDQ R14, R10
		ADCQ R15, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Compute a.x - a.y = BX||SI - R14||R15. Store in (R13*2^64+R12).
		MOVQ R14, R12
		MOVQ R15, R13

		NOTQ R12
		NOTQ R13
		BTRQ $63, R13

		ADDQ BX, R12
		ADCQ SI, R13
		BTRQ $63, R13
		ADCQ $0, R12
		ADCQ $0, R13

		// Mult (R11*2^64+R10) * (R13*2^64+R12). Store in (R9*2^64+R8).
		MOVQ $0, CX
		MOVQ R10, AX
		MULQ R12
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ R10, AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ R11, AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * a.y = BX||SI * R14||R15. Store in (R11*2^64+R10).
		MOVQ $0, CX
		MOVQ BX, AX
		MULQ R14
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ BX, AX
		MULQ R15
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ SI, AX
		MULQ R14
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ SI, AX
		MULQ R15
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Double (R11*2^64+R10) in-place.
		SHLQ $1, R11
		SHLQ $1, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10

	// E.Sub(E, A).Sub(E, B)
	MOVQ 0(SP), AX
	MOVQ 8(SP), BX
	MOVQ 16(SP), CX
	MOVQ 24(SP), DX

	MOVQ 32(SP), R12
	MOVQ 40(SP), R13
	MOVQ 48(SP), R14
	MOVQ 56(SP), R15

	NOTQ AX
	NOTQ BX
	BTRQ $63, BX
	NOTQ CX
	NOTQ DX
	BTRQ $63, DX
	NOTQ R12
	NOTQ R13
	BTRQ $63, R13
	NOTQ R14
	NOTQ R15
	BTRQ $63, R15

	ADDQ AX, R8
	ADCQ BX, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9
	ADDQ R12, R8
	ADCQ R13, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9
	ADDQ CX, R10
	ADCQ DX, R11
	BTRQ $63, R11
	ADCQ $0, R10
	ADCQ $0, R11
	ADDQ R14, R10
	ADCQ R15, R11
	BTRQ $63, R11
	ADCQ $0, R10
	ADCQ $0, R11

	// Current layout of memory:
	// AX, BX, CX, DX:     -A
	// R8, R9, R10, R11:    E
	// R12, R13, R14, R15: -B

	// Write E to stack in new location.
	MOVQ R8, 96(SP)
	MOVQ R9, 104(SP)
	MOVQ R10, 112(SP)
	MOVQ R11, 120(SP)

	// Current layout of stack:
	// A, B, C, E

	// Load B into R8...R11.
	MOVQ 32(SP), R8
	MOVQ 40(SP), R9
	MOVQ 48(SP), R10
	MOVQ 56(SP), R11

	// Calculate G := newGFp2().Sub(B, A) into R8...R11.
	ADDQ AX, R8
	ADCQ BX, R9
	BTRQ $63, R9
	ADCQ $0, R8
	ADCQ $0, R9

	ADDQ CX, R10
	ADCQ DX, R11
	BTRQ $63, R11
	ADCQ $0, R10
	ADCQ $0, R11

	// H := newGFp2().Add(B, A)
	// H.Neg(H)
	// Calculate H into AX...DX.
	ADDQ R12, AX
	ADCQ R13, BX
	BTRQ $63, BX
	ADCQ $0, AX
	ADCQ $0, BX

	ADDQ R14, CX
	ADCQ R15, DX
	BTRQ $63, DX
	ADCQ $0, CX
	ADCQ $0, DX

	MOVQ AX, 0(SP)
	MOVQ BX, 8(SP)
	MOVQ CX, 16(SP)
	MOVQ DX, 24(SP)
	MOVQ R8, 32(SP)
	MOVQ R9, 40(SP)
	MOVQ R10, 48(SP)
	MOVQ R11, 56(SP)

	// Load C into R12...R15
	MOVQ 64(SP), R12
	MOVQ 72(SP), R13
	MOVQ 80(SP), R14
	MOVQ 88(SP), R15

	// Calculate F := newGFp2().Sub(G, C) into R12...R15
	NOTQ R12
	NOTQ R13
	BTRQ $63, R13
	NOTQ R14
	NOTQ R15
	BTRQ $63, R15

	ADDQ R8, R12
	ADCQ R9, R13
	BTRQ $63, R13
	ADCQ $0, R12
	ADCQ $0, R13

	ADDQ R10, R14
	ADCQ R11, R15
	BTRQ $63, R15
	ADCQ $0, R14
	ADCQ $0, R15

	// Current layout of memory:
	// AX, BX, CX, DX:     H
	// R8, R9, R10, R11:   G
	// R12, R13, R14, R15: F
	//
	// Current layout of stack:
	// H, G, C, E

	// feMul(&c.x, E, F)
		// Mult a.y * b.y = 112(SP)||120(SP) * R14||R15. Store in (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 112(SP), AX
		MULQ R14
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ 112(SP), AX
		MULQ R15
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ R14
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ R15
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Negate (R9*2^64 + R8).
		NOTQ R8
		NOTQ R9
		BTRQ $63, R9

		// Mult a.x * b.x = 96(SP)||104(SP) * R12||R13. Add to (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 96(SP), AX
		MULQ R12
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		MOVQ 96(SP), AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX

		MOVQ 104(SP), AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * b.y = 96(SP)||104(SP) * R14||R15. Store in (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 96(SP), AX
		MULQ R14
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 96(SP), AX
		MULQ R15
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ R14
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ R15
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Mult a.y * b.x = 112(SP)||120(SP) * R12||R13. Add to (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 112(SP), AX
		MULQ R12
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX
		MOVQ 112(SP), AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Move out.
		MOVQ R8, 0(DI)
		MOVQ R9, 8(DI)
		MOVQ R10, 16(DI)
		MOVQ R11, 24(DI)

	// feMul(&c.y, G, H)
		// Mult a.y * b.y = 48(SP)||56(SP) * 16(SP)||24(SP). Store in (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 48(SP), AX
		MULQ 16(SP)
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ 48(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ 16(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Negate (R9*2^64 + R8).
		NOTQ R8
		NOTQ R9
		BTRQ $63, R9

		// Mult a.x * b.x = 32(SP)||40(SP) * 0(SP)||8(SP). Add to (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 32(SP), AX
		MULQ 0(SP)
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		MOVQ 32(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX

		MOVQ 40(SP), AX
		MULQ 0(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * b.y = 32(SP)||40(SP) * 16(SP)||24(SP). Store in (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 32(SP), AX
		MULQ 16(SP)
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 32(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ 16(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Mult a.y * b.x = 48(SP)||56(SP) * 0(SP)||8(SP). Add to (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 48(SP), AX
		MULQ 0(SP)
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX
		MOVQ 48(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ 0(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Move out.
		MOVQ R8, 32(DI)
		MOVQ R9, 40(DI)
		MOVQ R10, 48(DI)
		MOVQ R11, 56(DI)

	// feMul(&c.t, E, H)
		// Mult a.y * b.y = 112(SP)||120(SP) * 16(SP)||24(SP). Store in (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 112(SP), AX
		MULQ 16(SP)
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ 112(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ 16(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Negate (R9*2^64 + R8).
		NOTQ R8
		NOTQ R9
		BTRQ $63, R9

		// Mult a.x * b.x = 96(SP)||104(SP) * 0(SP)||8(SP). Add to (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 96(SP), AX
		MULQ 0(SP)
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		MOVQ 96(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX

		MOVQ 104(SP), AX
		MULQ 0(SP)
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * b.y = 96(SP)||104(SP) * 16(SP)||24(SP). Store in (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 96(SP), AX
		MULQ 16(SP)
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 96(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ 16(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 104(SP), AX
		MULQ 24(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Mult a.y * b.x = 112(SP)||120(SP) * 0(SP)||8(SP). Add to (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 112(SP), AX
		MULQ 0(SP)
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX
		MOVQ 112(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ 0(SP)
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 120(SP), AX
		MULQ 8(SP)
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Move out.
		MOVQ R8, 64(DI)
		MOVQ R9, 72(DI)
		MOVQ R10, 80(DI)
		MOVQ R11, 88(DI)

	// feMul(&c.z, F, G)
		// Mult a.y * b.y = 48(SP)||56(SP) * R14||R15. Store in (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 48(SP), AX
		MULQ R14
		MOVQ AX, R8
		MOVQ DX, R9
		MOVQ 48(SP), AX
		MULQ R15
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ R14
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ R15
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Negate (R9*2^64 + R8).
		NOTQ R8
		NOTQ R9
		BTRQ $63, R9

		// Mult a.x * b.x = 32(SP)||40(SP) * R12||R13. Add to (R9*2^64 + R8).
		MOVQ $0, CX
		MOVQ 32(SP), AX
		MULQ R12
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		MOVQ 32(SP), AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX

		MOVQ 40(SP), AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R8
		ADCQ AX, R9
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R8
		ADCQ DX, R9
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R9
		ADCQ CX, R8
		ADCQ $0, R9
		BTRQ $63, R9
		ADCQ $0, R8
		ADCQ $0, R9

		// Mult a.x * b.y = 32(SP)||40(SP) * R14||R15. Store in (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 32(SP), AX
		MULQ R14
		MOVQ AX, R10
		MOVQ DX, R11
		MOVQ 32(SP), AX
		MULQ R15
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ R14
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 40(SP), AX
		MULQ R15
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Mult a.y * b.x = 48(SP)||56(SP) * R12||R13. Add to (R11*2^64 + R10).
		MOVQ $0, CX
		MOVQ 48(SP), AX
		MULQ R12
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX
		MOVQ 48(SP), AX
		MULQ R13
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ R12
		SHLQ $1, DX
		ADDQ DX, R10
		ADCQ AX, R11
		ADCQ $0, CX
		MOVQ 56(SP), AX
		MULQ R13
		SHLQ $1, DX
		SHLQ $1, AX
		ADCQ $0, DX
		ADDQ AX, R10
		ADCQ DX, R11
		ADCQ $0, CX

		SHLQ $1, CX
		BTRQ $63, R11
		ADCQ CX, R10
		ADCQ $0, R11
		BTRQ $63, R11
		ADCQ $0, R10
		ADCQ $0, R11

		// Move out.
		MOVQ R8, 96(DI)
		MOVQ R9, 104(DI)
		MOVQ R10, 112(DI)
		MOVQ R11, 120(DI)

	RET
