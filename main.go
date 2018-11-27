package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Cat struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var cats = []Cat {
	Cat { ID: 1, Name:"Elsa", Age: 16 },
	Cat { ID: 2, Name:"Tijger", Age: 12 },
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, Mux!\n")
}

func CatsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
}

func CatHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	for _, cat := range cats {
		if cat.ID == id {
			json.NewEncoder(w).Encode(cat)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	log.Print("Server running at :8080")
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/cats", CatsHandler).Methods("GET")
	r.HandleFunc("/cats/{id:[0-9]+}", CatHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}