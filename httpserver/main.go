package main

import (
	"encoding/json"
    "fmt"
    "net/http"
)

type Product struct{
	Name	string	`json:"name"`
	Price 	float32	`json:"price"`
}

func products(w http.ResponseWriter, r *http.Request){
	items := []Product{
		{"Remera", 12.4},
		{"Short", 12.4},
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
	})
	
	http.HandleFunc("/products", products )
	
	http.ListenAndServe(":8080", nil)
}