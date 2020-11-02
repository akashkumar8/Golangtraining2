package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Heyyy folks...I'm from kloudone.Hope you doing well!")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
