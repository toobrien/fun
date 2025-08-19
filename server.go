package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world!")
}

func main() {

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
