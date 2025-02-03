package types

import (
	"fmt"
)

func DoTypes1() {
	var num int64
	fmt.Scan(&num)

	convert(num)
	fmt.Println(num)
}

func convert(x int64) uint16 {
	return uint16(x)
}
