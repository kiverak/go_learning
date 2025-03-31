package goroutines

import (
	"fmt"
	"time"
)

func DoGoroutines01() {
	fmt.Println("Обычный запуск")
	t := time.Now().Unix()
	for i := range 20 {
		myFunc(i)
	}

	time.Sleep(1 * time.Second) // Пауза в 1 секунду
	fmt.Println("Заняло времени:", time.Now().Unix()-t)

	fmt.Println("Параллельный запуск")

	t = time.Now().Unix()
	for i := range 20 {
		go myFunc(i)
	}

	time.Sleep(1 * time.Second) // Пауза в 1 секунду
	fmt.Println("Заняло времени:", time.Now().Unix()-t)

}

func myFunc(i int) {
	time.Sleep(1 * time.Second) // Пауза в 1 секунду
	fmt.Println(i)
}
