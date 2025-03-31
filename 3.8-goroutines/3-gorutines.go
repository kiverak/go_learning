package goroutines

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

func DoGoroutines3() {
	arr := []string{"apple", "Apple", "banana", "banana", "BANANA", "orange", "Orange", "orange"}
	inputStream := make(chan string, len(arr))
	outputStream := make(chan string, len(arr))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range arr {
			inputStream <- arr[rand.Intn(len(arr))]
		}
		close(inputStream)
	}()
	go func() {
		defer wg.Done()
		removeDuplicates(inputStream, outputStream)
	}()
	wg.Wait()
	for val := range outputStream {
		fmt.Println(val)
	}

}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var cur string
	for next := range inputStream {
		if strings.Compare(cur, next) != 0 {
			outputStream <- next
		}
		cur = next
	}
	close(outputStream)
}
