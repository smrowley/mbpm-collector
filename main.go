package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
