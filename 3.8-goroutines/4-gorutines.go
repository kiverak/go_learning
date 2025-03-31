package goroutines

import "fmt"

func DoGoroutines4() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		work()
	}()
	<-done
}

func work() {
	fmt.Println("Привет, work")
}
