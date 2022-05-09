package inspect

import (
    `fmt`
    `unsafe`
)

var (
    g_goid   = gfield("goid")
    g_labels = gfield("labels")
)

// Goroutine represents a goroutine object.
type Goroutine struct {
    _ uintptr
}

// ID returns the Goroutine ID.
func (self *Goroutine) ID() int {
    return int(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(self)) + g_goid)))
}

// String implements fmt.Stringer.
func (self *Goroutine) String() string {
    return fmt.Sprintf("<goroutine %d>", self.ID())
}

// Labels returns the current Goroutine labels.
func (self *Goroutine) Labels() map[string]string {
    return **(**map[string]string)(unsafe.Pointer(uintptr(unsafe.Pointer(self)) + g_labels))
}

func gfield(name string) uintptr {
    if t := FindType("runtime.g"); t == nil {
        panic("cannot find type 'runtime.g'")
    } else if fp, ok := t.FieldByName(name); !ok {
        panic("cannot get offset of '" + name + "' from 'runtime.g'")
    } else {
        return fp.Offset
    }
}
