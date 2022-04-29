// +build go1.17

package inspect

import (
    `os`
    `syscall`
    `unsafe`
)

const (
    _AP = syscall.MAP_ANON  | syscall.MAP_PRIVATE
    _RX = syscall.PROT_READ | syscall.PROT_EXEC
    _RW = syscall.PROT_READ | syscall.PROT_WRITE
)

var (
	gfunc = asmgfunc()
)

func mkptr(m uintptr) unsafe.Pointer {
    return *(*unsafe.Pointer)(unsafe.Pointer(&m))
}

func alignup(n uintptr, a int) uintptr {
    return (n + uintptr(a) - 1) &^ (uintptr(a) - 1)
}

func asmgfunc() func() *Goroutine {
    var mem uintptr
    var err syscall.Errno

    /* align the size to pages */
    nf := uintptr(len(gcode))
    nb := alignup(nf, os.Getpagesize())

    /* allocate a block of memory */
    if mem, _, err = syscall.Syscall6(syscall.SYS_MMAP, 0, nb, _RW, _AP, 0, 0); err != 0 {
        panic(err)
    }

    /* fill the code */
    buf := slice { mkptr(mem), int(nf), int(nf) }
    copy(*(*[]byte)(unsafe.Pointer(&buf)), gcode[:])

    /* protect the memory */
    if _, _, err = syscall.Syscall(syscall.SYS_MPROTECT, mem, nb, _RX); err != 0 {
        panic(err)
    }

    /* build the function */
    fp := unsafe.Pointer(&mem)
    return *(*func() *Goroutine)(unsafe.Pointer(&fp))
}

// G returns the current goroutine handle.
func G() *Goroutine {
    return gfunc()
}
