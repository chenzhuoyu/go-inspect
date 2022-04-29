package inspect

import (
    `reflect`
)

// FindType finds the type by name (such as "runtime.g") even if the type is
// not exported. This function returns nil if the type does not exist.
func FindType(name string) reflect.Type {
    var ok bool
    var rt reflect.Type

    /* attempt to find the type */
    EnumerateTypes(func(t reflect.Type) bool {
        t = dereftype(t)
        rt, ok = t, t.PkgPath() + "." + t.Name() == name
        return !ok
    })

    /* check if it exists */
    if ok {
        return rt
    } else {
        return nil
    }
}

// EnumerateTypes finds all the types in the current executable, whether it's
// exported or non-exported.
// If callback returns false, the enumeration stops.
func EnumerateTypes(callback func(reflect.Type) bool) {
    if ty, lr := typelinks(); len(ty) != 0 {
        for i, ln := range lr {
            for _, off := range ln {
                if !callback(typefrom(resolveTypeOff(ty[i], off))) {
                    break
                }
            }
        }
    }
}
