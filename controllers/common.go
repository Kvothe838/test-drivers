package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteStatus(response http.ResponseWriter, status int) {
	response.WriteHeader(status)
}

func SendJSONResponse(response http.ResponseWriter, status int, json interface{}) {
	response.Header().Set("Content-Type", "application/json")
	WriteStatus(response, status)
}

func Decode(body io.ReadCloser, target interface{}, response http.ResponseWriter) error {
	err := json.NewDecoder(body).Decode(target)
	if err != nil {
		fmt.Printf("error decoding %v. error: %v\n", body, err)
		http.Error(response, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
