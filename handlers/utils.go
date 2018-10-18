package handlers

import "net/http"

func sendError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(`{"message":"` + msg + `"}`)) // nolint: errcheck, gosec
}
