//this is go file
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Detail struct {
	Name          string `json:"Name"`
	Designation   string `json:"designation"`
	Qualification string `json:"qualification"`
}

var Details []Detail

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello this is my very first program in go and i will get used to it as quickly as possible")
	fmt.Println("API created successfully")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	Details = []Detail{
		Detail{Name: "Lakshmanan", Designation: "Software engineer trainee", Qualification: "B.E"},
		Detail{Name: "Venkat", Designation: "System admin", Qualification: "B.E"},
	}
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Details)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/details", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
