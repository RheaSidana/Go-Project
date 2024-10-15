package initializer

import "os"

func GetAppPort() string {
	return ":" + os.Getenv("APP_PORT")
}