package task8

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestIncrementCounter(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	goroutines := 100
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go IncrementCounter(&counter, &wg, &mu)
	}

	wg.Wait()

	assert.Equal(t, counter, goroutines, "counter mismatch")
}
