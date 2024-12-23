// Package server provides functionality for running the calculator in server mode.
package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"gocalc/pkg/calc"
)

const (
	entryPoint = "/api/v1/calculate"
)

var (
	// Handler for the calculation request
	Handler = withErrorHandling(calcHandler)
)

// Input represents the input data for the calculation
type Input struct {
	Expression string `json:"expression"`
}

// Output represents the result of the calculation
type Output struct {
	Result float64 `json:"result"`
}

// Error represents an error message
type Error struct {
	Error string `json:"error"`
}

// Return 500 status code and error message if panic occurred
func withErrorHandling(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Error{Error: err.(string)})
			}
		}()
		next(w, r)
	}
}

// calcHandler handles the calculation request
func calcHandler(w http.ResponseWriter, r *http.Request) {
	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Error{Error: err.Error()})
		return
	}

	if strings.Contains(input.Expression, "Python") {
		panic("Горфер паникует")
	}

	result, err := calc.Calc(input.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Error{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Output{Result: result})
}

// StartServer starts the HTTP server on the specified IP address and port
func StartServer(ipPort string) {
	http.HandleFunc(entryPoint, withErrorHandling(calcHandler))
	log.Printf("Server started on %s", ipPort)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error: %v", err)
		}
		log.Printf("Server stopped")
	}()
	err := http.ListenAndServe(ipPort, nil)
	if err != nil {
		panic(err)
	}
}
