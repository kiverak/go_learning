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
		return x * x
	}
	fmt.Println("\nСоздаю каналы, начитаю работу.")
	in1 := make(chan int, Num)
	in2 := make(chan int, Num)
	out := make(chan int, Num)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, Num)
	for i := 0; i < Num; i++ {
		in1 <- i
		in2 <- (i + Num)
		fmt.Printf("%v) В канал 1 послано: %v, В канал 2 послано: %v.\n", i+1, i, (i + Num))
	}
	fmt.Print("Функция fn с разной паузой возводит в квадрат значения из каналов.\n",
		"Итоговый результат fn(x1) + fn(x2)\n",
		"Барабанная дробь...\n")

	calcFail := false
	timeFail := false
	pass := "OK"
	for i := 0; i < Num; i++ {
		c := <-out
		result := i*i + (i+Num)*(i+Num)
		if c != result {
			calcFail = true
			pass = "Failed."
		}
		fmt.Printf("%v) Результат: %v. Правильный %v. -> %v\n", i+1, c, result, pass)
	}
	if calcFail {
		fmt.Println("Failed. Результат выполнения не соответствует ожидаемому.")
	}
	duration := time.Since(start)
	if duration.Seconds() > Num {
		timeFail = true
		fmt.Println("Failed. Время превышено.")
	}
	if !calcFail && !timeFail {
		fmt.Println("Passed. OK. Ты сделал это!\nВремя работы: ", duration)
	} else {
		fmt.Println("Пока неверно. Упорства вам не занимать, попробуете еще раз?\nВремя выполнения: ", duration)
	}
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		var wg sync.WaitGroup
		mapa := sync.Map{}
		for i := 0; i < n; i++ {
			mapa.Store(i, 0)
		}

		wg.Add(n * 2)

		for i := 0; i < n; i++ {
			x1 := <-in1
			calcAndStore(x1, fn, &mapa, i, &wg)
		}

		for i := 0; i < n; i++ {
			x2 := <-in2
			calcAndStore(x2, fn, &mapa, i, &wg)
		}

		wg.Wait()

		sendToOut(&mapa, out, n)
	}()
}

func calcAndStore(x int, fn func(int) int, mapa *sync.Map, i int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		fnRes := fn(x)
		res, _ := mapa.Load(i)
		val, _ := res.(int)
		mapa.Store(i, val+fnRes)
	}()
}

func sendToOut(s *sync.Map, out chan<- int, n int) {
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			res, _ := s.Load(i)
			if val, ok := res.(int); ok {
				out <- val
			}
		}
	}()
}
