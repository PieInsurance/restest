package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

// _prefixes we want to use.
var _prefixes = map[string]int{
	"T": 12,
	"G": 9,
	"M": 6,
	"":  0,
	"k": 3,
	"n": -9,
}

// errResponse returns a nice error/
func errResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add("Content-Type", "text/text")
	_, _ = w.Write([]byte(err.Error()))

	fmt.Printf("errored: %s\n", err)
}

// prefix encodes an SI prefix to the value and applies it to the unit.
func prefix(value float64, unit string) string {

	prefix := ""
	order := math.Log10(value)

	for pre, o := range _prefixes {
		if math.Abs(order-float64(_prefixes[prefix])) > math.Abs(order-float64(o)) {
			prefix = pre
		}
	}

	exp := _prefixes[prefix]
	short := value / math.Pow10(exp)

	return fmt.Sprintf("%.6E %s%s", short, prefix, unit)
}

// tempHandler decodes the path and returns an appropriate response.
func tempHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kelvinStr := vars["temp"]

	kelvin, err := strconv.ParseFloat(kelvinStr, 64)
	if err != nil {
		errResponse(w, err)
		return
	}

	fmt.Printf("running: %f\n", kelvin)

	celsius := kelvin - 273.15
	λ := 0.002898 / kelvin

	values := []string{
		prefix(kelvin, "K"),
		fmt.Sprintf("%e °C", celsius),
		prefix(λ, "m"),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(values)
}

// main entry point.
func main() {
	router := mux.NewRouter()
	router.HandleFunc(`/{temp}`, tempHandler)
	router.HandleFunc(`/{temp:[+-]?\d*\.?\d+}`, tempHandler)
	router.HandleFunc(`/{temp:[+-]?\d*\.?\d+[Ee][+-]?\d+}`, tempHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Print("started service\n")

	log.Fatal(srv.ListenAndServe())
}
