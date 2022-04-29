package inspect

import (
    `fmt`
    `reflect`
    `unsafe`
)


var (
    offset = uintptr(0)
)

// Goroutine represents a goroutine object.
type Goroutine struct {
    _ uintptr
}

// ID returns the goroutine ID.
func (self *Goroutine) ID() int {
    return int(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(self)) + offset)))
}

// String implements fmt.Stringer.
func (self *Goroutine) String() string {
    return fmt.Sprintf("<goroutine %d>", self.ID())
}

func init() {
    EnumerateTypes(func(t reflect.Type) bool {
        if t.Name() != "g" || t.PkgPath() != "runtime" {
            return true
        } else if fp, ok := t.FieldByName("goid"); ok {
            offset = fp.Offset
            return false
        } else {
            panic("cannot get offset of 'goid' from 'runtime.g'")
        }
    })
}
