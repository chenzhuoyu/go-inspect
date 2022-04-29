package inspect
import (
    `fmt`
    `reflect`
    `unsafe`
)

// FieldOps is the handle to a certain field.
type FieldOps struct {
    t *rtype
    p unsafe.Pointer
}

func newFieldOps(self reflect.Value, fv reflect.StructField) FieldOps {
    return FieldOps {
        t: unpacktype(fv.Type),
        p: unsafe.Pointer(uintptr((*eface)(unsafe.Pointer(&self)).ptr) + fv.Offset),
    }
}

func (self FieldOps) read() eface {
    if self.t.indirect() {
        return eface { typ: self.t, ptr: self.p }
    } else {
        return eface { typ: self.t, ptr: *(*unsafe.Pointer)(self.p) }
    }
}

func (self FieldOps) addr(v reflect.Value) unsafe.Pointer {
    if self.t.indirect() {
        return (*eface)(unsafe.Pointer(&v)).ptr
    } else {
        return unsafe.Pointer(&((*eface)(unsafe.Pointer(&v)).ptr))
    }
}

// Get returns the value of the referenced field, even if it's private.
func (self FieldOps) Get() reflect.Value {
    if self.t == nil {
        panic("inspect: invalid field")
    } else {
        return reflect.ValueOf(self.read().pack())
    }
}

// Set updates the value of the referenced field, even if it's private.
func (self FieldOps) Set(v reflect.Value) {
    if self.t == nil {
        panic("inspect: invalid field")
    } else {
        typedmemmove(self.t, self.p, self.addr(v))
    }
}

// FieldAt locates a field with index.
func FieldAt(self reflect.Value, idx int) (FieldOps, bool) {
    if idx < 0 || idx >= self.NumField() {
        return FieldOps{}, false
    } else {
        return newFieldOps(self, self.Type().Field(idx)), true
    }
}

// FieldByName locates a field with name.
func FieldByName(self reflect.Value, name string) (FieldOps, bool) {
    if fv, ok := self.Type().FieldByName(name); !ok {
        return FieldOps{}, false
    } else {
        return newFieldOps(self, fv), true
    }
}

// FieldAtOrPanic is like FieldAt, but it will panic if the index is invalid.
func FieldAtOrPanic(self reflect.Value, idx int) FieldOps {
    if fp, ok := FieldAt(self, idx); !ok {
        panic(fmt.Sprintf("inspect: field index out of range: %d", idx))
    } else {
        return fp
    }
}

// FieldByNameOrPanic is like FieldByName, but it will panic if the field does not exist.
func FieldByNameOrPanic(self reflect.Value, name string) FieldOps {
    if fp, ok := FieldByName(self, name); !ok {
        panic("inspect: no such field: " + name)
    } else {
        return fp
    }
}
