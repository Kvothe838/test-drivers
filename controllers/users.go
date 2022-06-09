package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kvothe838/drivers-api/services"
)

func Login(response http.ResponseWriter, request *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Profile  string `json:"profile"`
	}

	err := Decode(request.Body, &data, response)
	if err != nil {
		WriteStatus(response, http.StatusBadRequest)
		return
	}

	user, err := services.Login(data.Username, data.Password)
	if err != nil {
		if err == services.UserOrPasswordIncorrect {
			WriteStatus(response, http.StatusForbidden)
			return
		} else {
			fmt.Printf("error at login: %v", err)
			WriteStatus(response, http.StatusInternalServerError)
			return
		}
	}

	if user.Profile.Name != data.Profile {
		WriteStatus(response, http.StatusForbidden)
		return
	}

	/* token, err := GetToken(user)
	if err != nil {
		WriteStatus(response, http.StatusInternalServerError)
		return
	}

	request.AddCookie(&http.Cookie{
		Name:  "jwt-token",
		Value: *token,
	}) */

	request.AddCookie(&http.Cookie{
		Name:  "userId",
		Value: fmt.Sprintf("%d", user.Id),
	})

	WriteStatus(response, http.StatusOK)
}
