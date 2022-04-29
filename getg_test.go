package inspect

import (
    `testing`
)

func TestGetG(t *testing.T) {
    println(G().String())
}

func BenchmarkGetG_Serial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        G()
    }
}

func BenchmarkGetG_Parallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            G()
        }
    })
}
