package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		input          Input
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:           "valid expression",
			input:          Input{Expression: "2+2"},
			expectedStatus: http.StatusOK,
			expectedBody:   Output{Result: 4},
		},
		{
			name:           "invalid expression",
			input:          Input{Expression: "2/0"},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Error{Error: "деление на ноль"},
		},
		{
			name:           "invalid JSON",
			input:          Input{Expression: ""},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   Error{Error: "некорректное выражение"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req, err := http.NewRequest("POST", entryPoint, bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(calcHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			var expectedBody bytes.Buffer
			json.NewEncoder(&expectedBody).Encode(tt.expectedBody)
			if rr.Body.String() != expectedBody.String() {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expectedBody.String())
			}
		})
	}
}
