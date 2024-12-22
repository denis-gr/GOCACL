package main

import (
    "flag"
	"log"
	"net/http"
	"encoding/json"


	"gocalc/pkg/calc"
)

const (
	entryPoint = "/api/v1/calculate"
)

var (
	Handler = withErrorHandling(calcHandler)
)

type Input struct {
	Expression string `json:"expression"`
}

type Output struct {
	Result float64 `json:"result"`
}

type Error struct {
	Error string `json:"error"`
}


func withErrorHandling(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(Error{Error: "Internal Server Error"})
            }
        }()
        next(w, r)
    }
}


func calcHandler(w http.ResponseWriter, r *http.Request) {
	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Error{Error: err.Error()})
		return
	}

	result, err := calc.Calc(input.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Error{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Output{Result: result})
}


func startServer(ipPort string) {
	http.HandleFunc(entryPoint, withErrorHandling(calcHandler))
	log.Printf("Server started on %s", ipPort)
	http.ListenAndServe(ipPort, nil)
}

func main() {
	ipPort := flag.String("ipPort", ":8080", "IP:Port to listen on")
	flag.Parse()
	startServer(*ipPort)
}
