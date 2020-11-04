// Golang program to illustrate
// how to get a random number
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	res_1 := rand.Float64()
	res_2 := rand.Float64()
	res_3 := rand.Float64()

	fmt.Println("My First Random Number: ", res_1)
	fmt.Println("MySecond Random Number: ", res_2)
	fmt.Println("My Third Random Number: ", res_3)
}
