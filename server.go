package main

import (
	"fmt"
	"log"
	"net/http"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/gonum/stat/sampleuv"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}

func next_(w http.ResponseWriter, r *http.Request) {

	b := distuv.Bernoulli{P: 0.5}
	out := make([]float64, 1000)
	smpl := sampleuv.IIDer{Dist: b}
	smpl.Sample(out)

	fmt.Fprintf(w, "%v", out)

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/next_", next_)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
