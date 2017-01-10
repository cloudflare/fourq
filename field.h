#include "base.h"

#define feMov(a0,a1,a2,a3, c0,c1,c2,c3) \
	MOVQ a0, c0 \
	MOVQ a1, c1 \
	MOVQ a2, c2 \
	MOVQ a3, c3

#define feNeg(c0,c1,c2,c3) \
	bfeNeg(c0,c1) \
	bfeNeg(c2,c3)

#define feDbl(c0,c1,c2,c3) \
	bfeDbl(c0,c1) \
	bfeDbl(c2,c3)

#define feAdd(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeAdd(a0,a1, c0,c1) \
	bfeAdd(a2,a3, c2,c3)

#define feSub(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeSub(a0,a1, c0,c1) \
	bfeSub(a2,a3, c2,c3)

// feReverseSub negates c0:..:c3 and stores a0:..:a3 - c0:..:c3 in a0:..:a3.
#define feReverseSub(a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeReverseSub(a0,a1, c0,c1) \
	bfeReverseSub(a2,a3, c2,c3)

#define feMul(ra, rb, rc) \
	\ // T0 = a0 * b0, (r11, r10, r9, r8) <- [reg_p1_0-8] * [reg_p2_0-8]
	MOVQ 0+rb, DX \
	MULXQ 0+ra, R8, R9 \
	MULXQ 8+ra, R10, AX \
	\ // pushq r15
	\ // pushq r14
	ADDQ R10, R9 \
	MOVQ 8+rb, DX \
	MULXQ 8+ra, R10, R11 \
	\ // pushq r13
	ADCQ AX, R10 \
	\ // pushq r12
	MULXQ 0+ra, DX, AX \
	ADCQ $0, R11 \
	ADDQ DX, R9 \
	\
	\ // T1 = a1 * b1, (r15, r14, r13, r12) <- [reg_p1_16-24] * [reg_p2_16-24]
	MOVQ 16+rb, DX \
	MULXQ 16+ra, R12, R13 \
	ADCQ AX, R10 \
	MULXQ 24+ra, R14, AX \
	ADCQ $0, R11 \
	MOVQ 24+rb, DX \
	ADDQ R14, R13 \
	MULXQ 24+ra, R14, R15 \
	ADCQ AX, R14 \
	ADCQ $0, R15 \
	MULXQ 16+ra, DX, AX \
	ADDQ DX, R13 \
	ADCQ AX, R14 \
	ADCQ $0, R15 \
	\
	\ // c0 = T0 - T1 = a0*b0 - a1*b1
	XORQ AX, AX \
	SUBQ R12, R8 \
	SBBQ R13, R9 \
	SBBQ R14, R10 \
	SBBQ R15, R11 \
	\
	SHLQ $1, R10, R11 \
	SHLQ $1, R9, R10 \
	MOVQ 16+rb, DX \
	BTRQ $63, R9 \
	\
	\ // T0 = a0 * b1, (r15, r14, r13, r12) <- [reg_p1_0-8] * [reg_p2_16-24]
	MULXQ 0+ra, R12, R13 \
	BTRQ $63, R11 \ // Add prime if borrow=1
	SBBQ $0, R10 \
	SBBQ $0, R11 \
	MULXQ 8+ra, R14, AX \
	ADDQ R14, R13 \
	MOVQ 24+rb, DX \
	MULXQ 8+ra, R14, R15 \
	ADCQ AX, R14 \
	ADCQ $0, R15 \
	MULXQ 0+ra, DX, AX \
	ADDQ DX, R13 \
	ADCQ AX, R14 \
	ADCQ $0, R15 \
	\
	\ // Reducing and storing c0
	ADDQ R8, R10 \
	ADCQ R9, R11 \
	BTRQ $63, R11 \
	ADCQ $0, R10 \
	ADCQ $0, R11 \
	\
	\ // T1 = a1 * b0, (r12, r11, r10, r9) <- [reg_p1_16-24] * [reg_p2_0-8]
	MOVQ 0+rb, DX \
	MULXQ 16+ra, R8, R9 \
	MOVQ R10, 0+rc \
	MULXQ 24+ra, R10, AX \
	MOVQ R11, 8+rc \
	ADDQ R10, R9 \
	MOVQ 8+rb, DX \
	MULXQ 24+ra, R10, R11 \
	ADCQ AX, R10 \
	ADCQ $0, R11 \
	MULXQ 16+ra, DX, AX \
	ADDQ DX, R9 \
	ADCQ AX, R10 \
	ADCQ $0, R11 \
	\
	\ // c1 = T0 + T1 = a0*b1 + a1*b0
	ADDQ R12, R8 \
	\ // popq r12
	ADCQ R13, R9 \
	\ // popq r13
	ADCQ R14, R10 \
	\ // popq r14
	ADCQ R15, R11 \
	\
	\ // Reducing and storing c1
	SHLQ $1, R10, R11 \
	SHLQ $1, R9, R10 \
	BTRQ $63, R9 \
	BTRQ $63, R11 \
	ADCQ R10, R8 \
	ADCQ R11, R9 \
	BTRQ $63, R9 \
	\ // popq r15
	ADCQ $0, R8 \
	ADCQ $0, R9 \
	MOVQ R8, 16+rc \
	MOVQ R9, 24+rc

#define feSquare(ra, rc) \
	\ // t0 = (r9, r8) = a0 + a1, (rcx, r14) <- a1
	MOVQ 0+ra, R10 \
	\ // pushq r14
	MOVQ 16+ra, R14 \
	SUBQ R14, R10 \
	MOVQ 8+ra, R11 \
	MOVQ 24+ra, CX \
	SBBQ CX, R11 \
	\
	\ // pushq r13
	BTRQ $63, R11 \
	\ // pushq r12
	SBBQ $0, R10 \
	\
	\ // t1 = (r11, r10) = a0 - a1
	MOVQ R10, DX \
	MOVQ 0+ra, R8 \
	ADDQ R14, R8 \
	MOVQ 8+ra, R9 \
	ADCQ CX, R9 \
	\
	\ //  c0 = t0 * t1 = (a0 + a1)*(a0 - a1), (rcx, r14, r13, r12) <- (r9, r8) * (r11, r10)
	MULXQ R8, R12, R13 \
	SBBQ $0, R11 \
	MULXQ R9, R14, AX \
	MOVQ R11, DX \
	ADDQ R14, R13 \
	MULXQ R9, R14, CX \
	MOVQ 8+ra, R9 \
	ADCQ AX, R14 \
	ADCQ $0, CX \
	MULXQ R8, DX, AX \
	MOVQ 0+ra, R8 \
	ADDQ DX, R13 \
	ADCQ AX, R14 \
	ADCQ $0, CX \
	\
	\ // t2 = (r9, r8) = 2*a0
	ADDQ R8, R8 \
	ADCQ R9, R9 \
	\
	\ // Reducing and storing c0
	SHLQ $1, R14, CX \
	SHLQ $1, R13, R14 \
	BTRQ $63, R13 \
	BTRQ $63, CX \
	ADCQ R14, R12 \
	ADCQ CX, R13 \
	BTRQ $63, R13 \
	ADCQ $0, R12 \
	ADCQ $0, R13 \
	MOVQ R12, 0+rc \
	MOVQ R13, 8+rc \
	\
	\ //  c1 = 2a0 * a1, (rcx, r14, r11, r10) <- (r9, r8) * [reg_p1_16-24]
	MOVQ 16+ra, DX \
	MULXQ R8, R10, R11 \
	\ // popq r12
	MULXQ R9, R14, AX \
	\ // popq r13
	ADDQ R14, R11 \
	MOVQ 24+ra, DX \
	MULXQ R9, R14, CX \
	ADCQ AX, R14 \
	ADCQ $0, CX \
	MULXQ R8, DX, AX \
	ADDQ DX, R11 \
	ADCQ AX, R14 \
	ADCQ $0, CX \
	\
	\ // Reduce and store c1.
	SHLQ $1, R14, CX \
	SHLQ $1, R11, R14 \
	BTRQ $63, R11 \
	BTRQ $63, CX \
	ADCQ R14, R10 \
	ADCQ CX, R11 \
	BTRQ $63, R11 \
	\ // popq r14
	ADCQ $0, R10 \
	ADCQ $0, R11 \
	MOVQ R10, 16+rc \
	MOVQ R11, 24+rc
