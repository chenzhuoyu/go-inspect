package inspect

import (
    `fmt`
    `reflect`
    `testing`
)

type Struct1 struct {
    x *int
}

type Struct2 struct {
    i int
    v Struct1
}

func TestFields(t *testing.T) {
    a := &(&struct { v int }{ 0x5678 }).v
    b := &(&struct { v int }{ 0xaabbccdd }).v
    v := &Struct2 {
        i: 123,
        v: Struct1 {
            x: a,
        },
    }
    f := FieldByNameOrPanic(reflect.ValueOf(v).Elem(), "v")
    println(fmt.Sprintf("%#v", f.Get().Interface()))
    println(fmt.Sprintf("%#v", v.v))
    println("-------------")
    x := Struct1 { x: b }
    f.Set(reflect.ValueOf(x))
    println(fmt.Sprintf("%#v", x))
    println(fmt.Sprintf("%#v", f.Get().Interface()))
    println(fmt.Sprintf("%#v", v.v))
    println("-------------")
    println("a:", a, ", b:", b)
}
