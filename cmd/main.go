package main

import (
	"fmt"
	"io"
	"net/http"

	"mbpm-collector/pkg/domain"

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

	process := domain.NewProcess("SDLC")

	fmt.Printf("Process: %v\n", process)
}
