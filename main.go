package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kvothe838/drivers-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	post := BuildSetHandleFunc(router, "POST")

	post("/drivers", controllers.SaveDriver)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 100 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error while starting server: %s", err)
	} else {
		fmt.Println("api listening on port :8080")
	}
}

func BuildSetHandleFunc(router *mux.Router, method string) func(path string, handlerFunc http.HandlerFunc) {
	return func(path string, handlerFunc http.HandlerFunc) {
		router.HandleFunc(path, handlerFunc).Methods(method)
	}
}
