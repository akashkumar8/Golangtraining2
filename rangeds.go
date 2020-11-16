package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "] ")
	}
	fmt.Println("\nSum is :: ", sum)
	kvs := map[int]string{1: "apple", 2: "samsung"}
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}
	str := "Heyyy, Kloud!"
	for index, c := range str {
		fmt.Print("[", index, ",", string(c), "] ")
	}
}
