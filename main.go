package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	/*
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", mux))
	*/

	uuid, _ := uuid.NewRandom()

	workflow := Work{
		Description: uuid.String(),
	}

	fmt.Printf("Workflow: %v\n", workflow)
}
