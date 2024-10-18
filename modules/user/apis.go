package user

import "net/http"

var handlerRepository Repository
var userHandler Handler

func Apis() {
	handlerRepository = initRepository() 
	userHandler = initHandler(handlerRepository)

	http.HandleFunc("/users/", UserHandlerForGetOrDeleteOrPutOrPatch)
	http.HandleFunc("/users", UserHandlerForPostOrGetAllOrDeleteAll)
}

func UserHandlerForPostOrGetAllOrDeleteAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userHandler.GetAllUsersHandler(w, r)
	case "DELETE":
		userHandler.DeleteAllUsersHandler(w, r)
	case "POST":
		userHandler.CreateUserHandler(w, r)
	case "PUT":
		userHandler.PutUserHandler(w, r)
	case "PATCH":
		userHandler.PatchUserHandler(w, r)
	}

}
func UserHandlerForGetOrDeleteOrPutOrPatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		userHandler.GetUserHandler(w, r)
	case "DELETE":
		userHandler.DeleteUserHandler(w, r)
	}
}
