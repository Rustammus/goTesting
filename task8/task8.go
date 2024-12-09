package task8

import "sync"

func IncrementCounter(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*counter++
	mu.Unlock()
}
