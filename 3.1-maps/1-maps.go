package string

import (
	"fmt"
)

func DoMap1() {
	var arr [10]int
	cache := make(map[int]int, 10)

	for i := range arr {
		fmt.Scan(&arr[i])
	}

	for _, val := range arr {
		var tmp int
		cacheValue, exists := cache[val]
		if exists {
			tmp = cacheValue
		} else {
			tmp = work(val)
			cache[val] = tmp
		}

		fmt.Print(tmp, " ")
	}
}

func work(x int) int {
	return x * x
}
