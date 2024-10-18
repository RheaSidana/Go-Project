package user

import models "go-project/models"

type Response struct {
	User    models.User `json:"user"`
	Message string `json:"message"`
}

type ResponseList struct {
	Users []models.User `json:"users"`
	Message string `json:"message"`
}
