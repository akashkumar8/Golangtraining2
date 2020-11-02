package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bolType, _ := json.Marshal(false) //boolean Value
	fmt.Println(string(bolType))
	intType, _ := json.Marshal(15) // integer value
	fmt.Println(string(intType))
	fltType, _ := json.Marshal(1.677) //float value
	fmt.Println(string(fltType))
	strType, _ := json.Marshal("kloudgeeky") // string value
	fmt.Println(string(strType))
	slcA := []string{"akash", "kumar", "verma"} //slice value
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))
	mapA := map[string]int{"ak": 1, "kloudlearn": 2} //map value
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
