package goroutines

import (
	"fmt"
)

func DoGoroutines6() {
	first := make(chan int)
	second := make(chan int)
	stop := make(chan struct{})
	resultChan := calculator6(first, second, stop)

	// Отправляем значение в один из каналов (например, firstChan)
	first <- 5
	result := <-resultChan
	fmt.Println("Результат (1):", result) // Выведет: Результат: 25

	// Демонстрация с secondChan
	first2 := make(chan int)
	second2 := make(chan int)
	stop2 := make(chan struct{})
	resultChan2 := calculator6(first2, second2, stop2)

	second2 <- 10
	result2 := <-resultChan2
	fmt.Println("Результат (2):", result2) // Выведет: Результат: 30

	// Демонстрация с stopChan
	first3 := make(chan int)
	second3 := make(chan int)
	stop3 := make(chan struct{})
	resultChan3 := calculator6(first3, second3, stop3)

	stop3 <- struct{}{}
	// Попытка чтения из resultChan3 заблокируется, так как ничего не будет отправлено
	// В реальном приложении может потребоваться таймаут или другая логика обработки
	fmt.Println("Результат (stop):", <-resultChan3)
}

func calculator6(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close(outChan)
		select {
		case val := <-firstChan:
			outChan <- val * val
		case val := <-secondChan:
			outChan <- val * 3
		case <-stopChan:
			return
		}
	}()

	return outChan
}
