package inspect

import (
    `reflect`
    `unsafe`
)

// EnumerateTypes finds all the types in the current executable, whether it's
// exported or non-exported.
// If callback returns false, the enumeration stops.
func EnumerateTypes(callback func(reflect.Type) bool) {
    v0 := struct{}{}
    t0 := reflect.TypeOf(v0)
    types, links := typelinks()

    /* travel through all the classes */
    for i, link := range links {
        for _, class := range link {
            rt := resolveTypeOff(types[i], class)
            (*eface)(unsafe.Pointer(&t0)).ptr = rt

            /* only struct pointers */
            if t0.Kind() != reflect.Ptr || t0.Elem().Kind() != reflect.Struct {
                continue
            }

            /* get the struct type */
            tc := t0.Elem()
            tn := tc.Name()
            tp := tc.PkgPath()

            /* discard empty class names */
            if tn != "" && tp != "" {
                if !callback(tc) {
                    break
                }
            }
        }
    }
}
