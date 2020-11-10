package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"

)

type Item struct {
    UID string 'json:"UID"'
    Name string 'json:"Name"'
    Desc string 'json:"Desc"'
    Price float64 'json:"Price"'
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Endpoint called: homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Function Called: getInventory()")

    json.NewEncoder(w).Encode(inventory)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage).Methods("GET")
    router.HandleFunc("/inventory", getInventory).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))

    
}

func main() {
    inventory = append(inventory, Item{
        UID: "0",
        Name: "Cheese",
        Desc: "A fine blockof cheese",
        Price: 4.99,
	})
	inventory = append(inventory, Item{
        UID: "1",
        Name: "Cadburry",
        Desc: "A bunch of cadburry",
        Price: 3.25,
	})
	

    handleRequests()
}