package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/gonum/stat/sampleuv"
)

const (
	sample_size = 1000
	easy        = 0.60
	medium      = 0.55
	hard        = 0.51
)

var difficulty = map[string]float64{
	"easy":   0.60,
	"medium": 0.55,
	"hard":   0.51,
}

//var requests = map[string]bool{}

func index(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}

func next(w http.ResponseWriter, r *http.Request) {

	dif := r.URL.Query().Get("difficulty")

	if dif == "" {
		fmt.Fprintf(w, "missing querystring parameter difficulty")
		return
	}

	p_, ok := difficulty[dif]

	if !ok {
		fmt.Fprintf(w, "invalid difficulty: %v", dif)
		return
	}

	b := distuv.Bernoulli{P: p_}
	x := make([]float64, sample_size)
	smpl := sampleuv.IIDer{Dist: b}
	smpl.Sample(x)

	res := struct {
		Result []float64 `json:"res"`
	}{Result: x}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/next_", next)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
