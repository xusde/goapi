package api

import (
	"encoding/json"
	"net/http"
)

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	Code int
	Balance float64
}

// error response if error occurs
type ErrorResponse struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter,  message string, code int) {
	var errorResponse = ErrorResponse{
		Code: code,
		Message: message,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errorResponse)

}

var(
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
	}

)