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

func SendJSONResponse(response http.ResponseWriter, status int, jsonObj interface{}) {
	bytes, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Printf("error while mashalling object %v. err: %v", jsonObj, err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	WriteStatus(response, status)
	_, err = response.Write(bytes)
	if err != nil {
		fmt.Printf("error while writing json bytes: %v", err)
	}
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
