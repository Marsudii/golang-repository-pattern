package services

import (
	"go-fiber/models"
	"go-fiber/request"
	"go-fiber/response"
)

type UserService interface {
	GetAllData() ([]models.User, error)
	GetDetailData(param string) ([]models.User, error)
	New(userRequest request.UserRequest) (response.UserResponse, error)
	Validation(userRequest request.UserRequest) error
}
