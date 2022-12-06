package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)

	if body == nil || statusCode == http.StatusNoContent {
		return
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Unable to marshal response. Error %s", err)
		return
	}

	_, e := w.Write(bytes)
	if e != nil {
		log.Printf("Unable to write response. Error %s", e)
		return
	}
}
