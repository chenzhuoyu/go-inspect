// +build go1.16,!go1.17

package inspect

//go:nosplit
// G returns the current goroutine handle.
func G() *Goroutine
