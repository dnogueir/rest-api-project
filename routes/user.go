package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("getting users")
}

func GetUser(w http.ResponseWriter, rep *http.Request, p httprouter.Params) {
	fmt.Println("getting user by id")
}

func CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("Creating an user")
}

func UpdateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("updating user by id")
}

func DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("deleting user by id")
}
