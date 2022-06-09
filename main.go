package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kvothe838/drivers-api/controllers"
	"github.com/Kvothe838/drivers-api/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDatabase()
	router := mux.NewRouter()

	router.Use()

	post := BuildSetHandleFunc(router, "POST")
	get := BuildSetHandleFunc(router, "GET")

	post("/drivers", controllers.SaveDriver)
	get("/drivers", controllers.GetDrivers)

	post("/login", controllers.Login)

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
		/* handlerFuncThoughMiddleware := controllers.Middleware(handlerFunc) */
		router.HandleFunc(path, handlerFunc).Methods(method)
	}
}
