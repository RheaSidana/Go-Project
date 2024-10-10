package main

import (
	"fmt"
	"go-project/initializer"
	"os"
)

func init(){
	initializer.LoadEnvVariables()
}

func main() {
	port := os.Getenv("APP_PORT")

	fmt.Println("Port is : ", port)
}