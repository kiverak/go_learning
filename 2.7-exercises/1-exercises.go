package exercises2

import (
	"fmt"
	"math"
)

func DoExercise1() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(getРypotenuse(a, b))
}

func getРypotenuse(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
