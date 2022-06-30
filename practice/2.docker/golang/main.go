package main

import (
	"fmt"
	"log"
	"net/http"
    

	"github.com/gorilla/mux"
)

func HelloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Word!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloWord)
	log.Fatal(http.ListenAndServe(":8080", router))
}
