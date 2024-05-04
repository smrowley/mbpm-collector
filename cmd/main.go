package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"mbpm-collector/internal/domain"
)

var workflow domain.Workflow

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, fmt.Sprintf("Workflow: %v\n", workflow.GetId()))
}

func main() {
	workflow = domain.NewWorkflow("SDLC")

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
