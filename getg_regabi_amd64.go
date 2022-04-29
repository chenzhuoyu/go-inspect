// +build go1.17

package inspect

var gcode = [...]byte {
    0x4c, 0x89, 0xf0,   // MOVQ %r14, %rax
    0xc3,               // RET
}
