package inspect

import (
    `reflect`
    `testing`
)

func TestTypes(t *testing.T) {
    EnumerateTypes(func(t reflect.Type) bool {
        if t.PkgPath() == "" || t.Name() == "" {
            println(t.String())
        } else {
            println(t.PkgPath() + "." + t.Name())
        }
        return true
    })
}
