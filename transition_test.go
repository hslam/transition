package transition

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestTransition(t *testing.T) {
	lows := []uint64{1, 2, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 1, 1, 1, 1, 0, 0, 0, 0}
	highs := []uint64{0, 0, 0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 21, 22, 24, 25, 26, 27, 29, 30, 31, 32}
	for c := 1; c <= 32; c++ {
		wg := sync.WaitGroup{}
		base := uint64(1000)
		concurrency := func() int { return c }
		tran := NewTransition(4, concurrency)
		var l, h uint64
		var lock sync.Mutex
		var count uint64
		for i := 0; i < c; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for atomic.AddUint64(&count, 1) <= uint64(c)*base {
					low := func() {
						atomic.AddUint64(&l, 1)
					}
					high := func() {
						atomic.AddUint64(&h, 1)
					}
					lock.Lock()
					tran.Smooth(low, high)
					lock.Unlock()
				}
			}()
		}
		wg.Wait()
		if l != lows[c-1]*base {
			t.Errorf("expect %d, got %d", lows[c-1]*base, l)
		} else if h != highs[c-1]*base {
			t.Errorf("expect %d, got %d", highs[c-1]*base, h)
		}
	}
}
