package main

import (
	"fmt"
	"math/rand"
)

func floatrandom(value_1, value_2 float64) float64 {
	return value_1 + value_2 + rand.Float64()
}

func main() {

	res1 := floatrandom(28, 32)
	res2 := floatrandom(50, 70)
	res3 := floatrandom(100, 700)

	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
}
