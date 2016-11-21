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

#define feMul(carry, a0,a1,a2,a3, b0,b1,b2,b3, c0,c1,c2,c3) \
	bfeMul(carry, a2,a3, b2,b3, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	bfeNeg(c0,c1) \
	MOVQ $0, carry \
	bfeMulAdd(carry, a0,a1, b0,b1, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	\
	bfeMul(carry, a0,a1, b2,b3, c2,c3) \
	bfeMulAdd(CX, a2,a3, b0,b1, c2,c3) \
	bfeMulReduce(carry, c2,c3)

// feSquare squares a0:..:a3 and stores the result in c0:..:c3, using
// carry,t0,t1 as workspace.
#define feSquare(carry,t0,t1, a0,a1,a2,a3, c0,c1,c2,c3) \
	bfeMov(a0,a1, c2,c3) \
	bfeAdd(a2,a3, c2,c3) \
	\
	bfeMov(a2,a3, t0,t1) \
	bfeSub(a0,a1, t0,t1) \
	\
	bfeMul(carry, c2,c3, t0,t1, c0,c1) \
	bfeMulReduce(carry, c0,c1) \
	\
	bfeMul(carry, a0,a1, a2,a3, c2,c3) \
	bfeMulReduce(carry, c2,c3) \
	bfeDbl(c2,c3)
