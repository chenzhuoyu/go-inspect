// +build go1.16

package inspect

var gcode = [...]byte{
    0x80, 0x03, 0x00, 0x91, // mov  x0, x28
    0xc0, 0x03, 0x5f, 0xd6, // ret
}
