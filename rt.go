package inspect

import (
    `reflect`
    `unsafe`
)

const (
    _F_direct = 1 << 5
)

type slice struct {
    ptr unsafe.Pointer
    len int
    cap int
}

type eface struct {
    typ *rtype
    ptr unsafe.Pointer
}

func (self eface) pack() (r interface{}) {
    *(*eface)(unsafe.Pointer(&r)) = self
    return
}

type rtype struct {
    size       uintptr
    ptrdata    uintptr
    hash       uint32
    tflag      uint8
    align      uint8
    fieldAlign uint8
    kflags     uint8
    equal      func(unsafe.Pointer, unsafe.Pointer) bool
    gcdata     *byte
    str        int32
    ptrToThis  int32
}

func unpacktype(t reflect.Type) *rtype {
    return (*rtype)((*eface)(unsafe.Pointer(&t)).ptr)
}

func (self *rtype) indirect() bool {
    return (self.kflags & _F_direct) == 0
}

//go:noescape
//go:linkname typelinks reflect.typelinks
func typelinks() ([]unsafe.Pointer, [][]int32)

//go:noescape
//go:linkname typedmemmove reflect.typedmemmove
func typedmemmove(_ *rtype, _ unsafe.Pointer, _ unsafe.Pointer)

//go:noescape
//go:linkname resolveTypeOff runtime.resolveTypeOff
func resolveTypeOff(_ unsafe.Pointer, _ int32) unsafe.Pointer
