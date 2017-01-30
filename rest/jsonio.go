package rest

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/palantir/stacktrace"
	"io/ioutil"
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

func ReadBody(r *http.Request) ([]byte, error) {
	if r == nil || r.Body == nil {
		return nil, stacktrace.NewError("request bidy empty")
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot read input")
	}
	return data, nil
}