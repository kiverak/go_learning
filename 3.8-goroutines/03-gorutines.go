package goroutines

import (
	"fmt"
)

func DoGoroutines03() {
	<-myFunc03()
}

func myFunc03() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}
