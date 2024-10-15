package main

import (
	"go-project/initializer"
	"log"
	"net/http"
	"go-project/modules/user"
)

var port string

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	port = initializer.GetAppPort()
}

func main() {
	defer initializer.CloseDb()

	user.Apis()

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
