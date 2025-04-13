package goroutines

import (
	"fmt"
	"time"
)

func DoGoroutines7() {
	args := make(chan int)
	done := make(chan struct{})
	result := calculator7(args, done)

	go func() {
		args <- 1
		time.Sleep(100 * time.Millisecond)
		args <- 2
		time.Sleep(100 * time.Millisecond)
		args <- 3
		time.Sleep(100 * time.Millisecond)
		close(done) // Сигнал завершения
	}()

	finalSum := <-result
	fmt.Println("Сумма:", finalSum) // Выведет: Сумма: 6
}

func calculator7(arguments <-chan int, done <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		sum := 0
		for {
			select {
			case arg := <-arguments:
				sum += arg
			case <-done:
				resultChan <- sum
				return
			}
		}
	}()

	return resultChan
}
