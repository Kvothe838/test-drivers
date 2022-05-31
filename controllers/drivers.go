package controllers

import (
	"net/http"

	"github.com/Kvothe838/drivers-api/model"
	"github.com/Kvothe838/drivers-api/services"
)

func SaveDriver(response http.ResponseWriter, request *http.Request) {
	var driver model.Driver

	err := Decode(request.Body, &driver, response)
	if err != nil {
		return
	}

	services.SaveDriver(driver)
	WriteStatus(response, http.StatusNoContent)
}
