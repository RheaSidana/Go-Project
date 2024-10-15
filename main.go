package main

import (
	"go-project/initializer"
	"log"
	"net/http"
)

var port string

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	port = initializer.GetAppPort()
}

func main() {
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
