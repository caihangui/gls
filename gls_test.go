package gls

import (
	"sync"
	"testing"
)

func Benchmark(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(val int) {
			defer wg.Done()
			Store("val", val)
			func() {
				_, ok := Load("val")
				if !ok {
					return
				}
			}()
			Delete("val")
		}(i)
	}
	wg.Wait()
}
