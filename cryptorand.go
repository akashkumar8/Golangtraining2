package main 
  
import ( 
    "crypto/rand"
    "fmt"
) 
  
func main() { 
    RandomCrypto, _ := rand.Prime(rand.Reader, 128) 
    fmt.Println(RandomCrypto) 
}