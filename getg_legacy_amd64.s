// +build !go1.17

#include "go_asm.h"
#include "textflag.h"

TEXT ·G(SB), NOSPLIT, $0 - 8
    MOVQ    TLS, CX
    MOVQ    0(CX)(TLS), AX
    MOVQ    AX, ret+0(FP)
    RET
