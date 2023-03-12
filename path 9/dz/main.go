package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "send get request like /?radius=2")
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, reflect.TypeOf(r.URL.Query().Get("radius")))
		// fmt.Fprint(w, r.URL.Query().Get("radius") == "")
		val, err := strconv.ParseFloat(r.URL.Query().Get("radius"), 64)

		if err != nil {
			fmt.Fprint(w, "send get request like /?radius=2")
		} else {
			fmt.Fprint(w, getS(val))
		}
	})

	http.ListenAndServe(":80", nil)
}

func getS(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
