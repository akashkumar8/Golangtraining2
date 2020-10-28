package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "Heyyyy Akash"
	ch <- " shoutout from kloudone"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}