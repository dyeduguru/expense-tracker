package rest

import (
	"net/http"
	"encoding/json"
	"log"
)

func WriteJSON(w http.ResponseWriter, obj interface{}, status int) {
	if err, ok := obj.(error); ok {
		obj = err.Error()
	}
	data, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("content-Type", "applucation/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("Failed to write %v to the caller", string(data))
	}
}
