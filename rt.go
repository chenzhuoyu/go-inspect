package inspect

import (
    `reflect`
    `unsafe`
)

const (
    _F_direct = 1 << 5
)

var (
    itab_rtype = itabof(new(reflect.Type))
)

type itab struct {
    inter *rtype
    typ   *rtype
    hash  uint32
    _     [4]byte
    fn    [1]uintptr
}

type slice struct {
    ptr unsafe.Pointer
    len int
    cap int
}

type iface struct {
    tab *itab
    ptr unsafe.Pointer
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
func resolveTypeOff(_ unsafe.Pointer, _ int32) *rtype

func itabof(v interface{}) *itab {
    t := reflect.TypeOf(v).Elem()
    return (*iface)(unsafe.Pointer(&t)).tab
}

func typefrom(t *rtype) (r reflect.Type) {
    (*iface)(unsafe.Pointer(&r)).tab = itab_rtype
    (*iface)(unsafe.Pointer(&r)).ptr = unsafe.Pointer(t)
    return
}

func dereftype(t reflect.Type) reflect.Type {
    for t.Kind() == reflect.Ptr { t = t.Elem() }
    return t
}
