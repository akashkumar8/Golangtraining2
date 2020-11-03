package main  
import (  
    "sort"  
    "fmt"  
)  
func main() {  
  
    intValue := []int{3, 12, 43, 32, 23}  
    sort.Ints(intValue)  
    fmt.Println("Ints:   ", intValue)  
  
    floatValue := []float64{10.4, 4.5, 5.8, 9,5}  
    sort.Float64s(floatValue)  
    fmt.Println("floatValue:   ", floatValue)  
  
    stringValue := []string{"Akash", "Kumar", "Verma"}  
    sort.Strings(stringValue)  
    fmt.Println("Strings:", stringValue)  
  
    str := sort.Float64sAreSorted(floatValue)  
    fmt.Println("Sorted: ", str)  
} 