package controllers

import (
	"net/http"
)

func SaveDriver(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNoContent)
}
