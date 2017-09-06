package go_commons

import "sync"

var wg sync.WaitGroup

func ConcurrentFn(n int, fn func()) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}
