package inspect

import (
    `fmt`
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
    if t := FindType("runtime.g"); t == nil {
        panic("cannot find type 'runtime.g'")
    } else if fp, ok := t.FieldByName("goid"); !ok {
        panic("cannot get offset of 'goid' from 'runtime.g'")
    } else {
        offset = fp.Offset
    }
}
