package goroutines

import (
	"fmt"
)

func DoGoroutines02() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go myFunc1(c)
	}
	fmt.Println(<-c)

}

func myFunc1(c chan int) {
	c <- cap(c)
}
