package inspect

import (
    `reflect`
    `testing`
)

func TestTypes(t *testing.T) {
    EnumerateTypes(func(t reflect.Type) bool {
        println(t.String())
        return true
    })
}
