package main  
import "sort"  
import "fmt"  
  
type  OrderByLengthDesc []string  
func (s OrderByLengthDesc) Len() int {  
    return len(s)  
}  
func (str OrderByLengthDesc) Swap(i, j int) {  
    str[i], str[j] = str[j], str[i]  
}  
func (s OrderByLengthDesc) Less(i, j int) bool {  
    return len(s[i]) > len(s[j])  
}  
func main() {  
    city := []string{"Patna", "Bihar","Jharkhand","uttarpradesh"}  
    sort.Sort(OrderByLengthDesc(city))  
    fmt.Println(city)  
} 