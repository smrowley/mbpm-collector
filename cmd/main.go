package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"mbpm-collector/internal/domain"
)

var process domain.Process

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, fmt.Sprintf("Workflow: %v\n", process.GetId()))
}

func createInstance(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	io.WriteString(w, "test")
}

func main() {
	log.Println("Starting server")
	process = domain.NewProcess("SDLC")

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/createInstance", createInstance)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
