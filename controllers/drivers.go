package controllers

import (
	"net/http"

	"github.com/Kvothe838/drivers-api/model"
	"github.com/Kvothe838/drivers-api/services"
)

func SaveDriver(response http.ResponseWriter, request *http.Request) {
	var data struct {
		DNI      string `json:"DNI"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := Decode(request.Body, &data, response)
	if err != nil {
		return
	}

	newUser, err := services.SignUp(data.Username, data.Password)
	if err != nil {
		if err == services.UserAlreadyExists {
			WriteStatus(response, http.StatusConflict)
		} else {
			WriteStatus(response, http.StatusInternalServerError)
		}
	}

	driver := model.Driver{
		User:    *newUser,
		DNI:     data.DNI,
		Name:    data.Name,
		Surname: data.Surname,
	}

	services.SaveDriver(driver)
	WriteStatus(response, http.StatusOK)
}
