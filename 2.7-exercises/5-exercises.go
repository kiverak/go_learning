package exercises2

import (
	"fmt"
	"math"
)

var k, v, p float64

func DoExercise5() {
	fmt.Println(T())
}

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
