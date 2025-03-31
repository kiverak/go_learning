package goroutines

import (
	"fmt"
	"sync"
)

func DoGoroutines5() {
	var x int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			work5()
			mu.Lock()
			x++
			mu.Unlock()
		}(&wg, &mu)
	}
	wg.Wait()
	fmt.Println("x = ", x)
}

func work5() {
	fmt.Println("Привет, work")
}
