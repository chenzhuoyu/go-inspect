package inspect

import (
    `context`
    `fmt`
    `runtime/pprof`
    `testing`
)

func TestGetG(t *testing.T) {
    println(G().String())
}

func TestGLabels(t *testing.T) {
    pprof.Do(context.Background(), pprof.Labels("key1", "value1", "key2", "value2"), func(_ context.Context) {
        println(fmt.Sprintf("%#v", G().Labels()))
    })
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
