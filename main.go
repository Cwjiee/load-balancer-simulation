package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct{
	Name string `json:"name"`
	Description string `json:"description"`
}

var mockItems = []Item{
	{
		Name: "item 1",
		Description: "description 1",
	},
	{
		Name: "item 2",
		Description: "description 1",
	},
	{
		Name: "item 3",
		Description: "description 1",
	},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world")
	}).Methods("GET")

	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockItems)
		log.Println("response returned")
	}).Methods("GET")

	r.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal("wrong id")
			return
		}

		json.NewEncoder(w).Encode(mockItems[id])
		log.Println("response returned")
	}).Methods("GET")

	log.Println("starting server at port http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
