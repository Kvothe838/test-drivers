package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kvothe838/drivers-api/model"
	"github.com/Kvothe838/drivers-api/services"
	"github.com/gorilla/mux"
)

func SaveDriver(response http.ResponseWriter, request *http.Request) {
	isAuthorized, err := IsAuthorized(request, "save-drivers")
	if err != nil {
		WriteStatus(response, http.StatusInternalServerError)
		return
	}

	if !*isAuthorized {
		WriteStatus(response, http.StatusForbidden)
		return
	}

	var data struct {
		DNI      string `json:"DNI"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err = Decode(request.Body, &data, response)
	if err != nil {
		return
	}

	newUser, err := services.SignUp(data.Username, data.Password)
	if err != nil {
		if err == services.UserAlreadyExists {
			WriteStatus(response, http.StatusConflict)
			return
		} else {
			WriteStatus(response, http.StatusInternalServerError)
			return
		}
	}

	driver := model.Driver{
		User:    *newUser,
		DNI:     data.DNI,
		Name:    data.Name,
		Surname: data.Surname,
	}

	err = services.SaveDriver(driver)
	if err != nil {
		WriteStatus(response, http.StatusInternalServerError)
		return
	}

	WriteStatus(response, http.StatusOK)
}

func GetDrivers(response http.ResponseWriter, request *http.Request) {
	isAuthorized, err := IsAuthorized(request, "get-drivers")
	if err != nil {
		WriteStatus(response, http.StatusBadRequest)
		return
	}

	if !*isAuthorized {
		WriteStatus(response, http.StatusForbidden)
		return
	}

	page := mux.Vars(request)["pages"]

	parsedPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		WriteStatus(response, http.StatusBadRequest)
		return
	}

	drivers, err := services.GetDrivers(int(parsedPage))
	if err != nil {
		fmt.Printf("error getting drivers: %v\n", err)
		WriteStatus(response, http.StatusInternalServerError)
	}

	SendJSONResponse(response, http.StatusOK, drivers)
}

func GetNonTravellingDrivers(response http.ResponseWriter, request *http.Request) {
	isAuthorized, err := IsAuthorized(request, "get-drivers")
	if err != nil {
		WriteStatus(response, http.StatusBadRequest)
		return
	}

	if !*isAuthorized {
		WriteStatus(response, http.StatusForbidden)
		return
	}

	drivers, err := services.GetNonTravellingDrivers()
	if err != nil {
		fmt.Printf("error getting drivers: %v\n", err)
		WriteStatus(response, http.StatusInternalServerError)
	}

	SendJSONResponse(response, http.StatusOK, drivers)
}
