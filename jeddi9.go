package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	/*
	   Fix the race condition you created in the previous exercise by using a mutex
	   ● it makes sense to remove runtime.Gosched()

	*/
	//counter := 0
	var counter int64
	gr := 100

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			/*
				Fix the race condition you created in the previous exercise by using a mutex
					● it makes sense to remove runtime.Gosched()
			*/
			//mu.Lock()

			/*
				Fix the race condition you created in the previous exercise by using package atomic
			*/

			atomic.AddInt64(&counter, 1)

			v := counter
			//runtime.Gosched()
			v++
			//counter = v
			fmt.Println(counter)

			//	mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("end value", counter)
}
