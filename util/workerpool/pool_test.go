package workerpool

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func init() {
	println("using MAXPROC")
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(4, 10000)
	defer pool.StopAll()

	iterations := 1000000
	var counter uint64

	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		arg := uint64(1)
		job := func() {
			defer wg.Done()
			atomic.AddUint64(&counter, arg)
		}

		pool.JobQueue <- Job{F: job, Idx: i}
	}
	wg.Wait()

	counterFinal := atomic.LoadUint64(&counter)
	if uint64(iterations) != counterFinal {
		t.Errorf("iterations %v is not equal counterFinal %v", iterations, counterFinal)
	}
}
