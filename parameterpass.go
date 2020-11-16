package main

import "fmt"

func IncrementPassByValue(x int) {
	x++
}
func main1() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)
	IncrementPassByValue(i)
	fmt.Println("Value of i after increment is : ", i)
}
