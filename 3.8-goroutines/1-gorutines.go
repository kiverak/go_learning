package goroutines

import (
	"fmt"
)

func DoGoroutines1() {
	c := make(chan int, 5)
	go task(c, 100)
	fmt.Println(<-c)
}

func task(c chan int, n int) {
	c <- n + 1
}
