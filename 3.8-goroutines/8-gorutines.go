package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const Num = 3

func DoGoroutines8() {
	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(Num)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, Num)
	in2 := make(chan int, Num)
	out := make(chan int, Num)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, Num+1)
	for i := 0; i < Num+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < Num; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > Num {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	map1, map2 := make(map[int]int, n), make(map[int]int, n)
	var wg sync.WaitGroup
	wg.Add(2 * n)
	for i := 0; i < n; i++ {
		go func() {
			map1[i] = <-in1
			map2[i] = <-in2
		}()
	}

	for i := 0; i < n; i++ {
		go calc(map1, fn, i, &wg)
		go calc(map2, fn, i, &wg)
	}
	wg.Wait()

	go func() {
		for i := 0; i < n; i++ {
			out <- map1[i] + map2[i]
		}
	}()
}

func calc(map1 map[int]int, fn func(int) int, i int, wg *sync.WaitGroup) {
	go func(i int) {
		defer wg.Done()
		map1[i] = fn(map1[i])
	}(i)
}
