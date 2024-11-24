package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := map[string]interface{}{
			"success": false,
			"error":   "Method not allowed.",
		}

		responseBuffer, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		_, _ = w.Write(responseBuffer)
		return
	}

	_, bodyErr := io.ReadAll(r.Body)

	if bodyErr != nil {
		response := map[string]interface{}{
			"success": false,
			"error":   bodyErr.Error(),
		}

		responseBuffer, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write(responseBuffer)
		return
	}

	response := map[string]interface{}{
		"success": true,
	}

	responseBuffer, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(responseBuffer)
}

func main() {
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
