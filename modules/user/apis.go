package user

import "net/http"

func Apis() {
	repository := initRepository()
	userHandler := initHandler(repository)

	http.HandleFunc("/users", userHandler.CreateUserHandler)
}