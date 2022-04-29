// +build !go1.17

package inspect

//go:nosplit
// G returns the current goroutine handle.
func G() *Goroutine
