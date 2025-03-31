package goroutines

import (
	"fmt"
)

func DoGoroutines2() {
	c := make(chan string, 5)
	go task2(c, "100")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}
