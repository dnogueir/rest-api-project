package main

import (
	"net/http"
	r "rest-api-project/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/users", r.GetUsers)
	mux.GET("/user/:id", r.GetUser)
	mux.POST("/user", r.CreateUser)
	mux.PUT("/user/:id", r.UpdateUser)
	mux.DELETE("/user/:id", r.DeleteUser)
	http.ListenAndServe(":8080", mux)
}
