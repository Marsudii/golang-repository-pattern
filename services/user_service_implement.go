package services

import (
	"errors"
	"go-fiber/models"
	"go-fiber/repositories"
	"go-fiber/request"
	"go-fiber/response"
	"log"
)

type UserServiceImplement struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) UserService {
	return &UserServiceImplement{
		userRepo: *userRepo,
	}
}

// GetAllData implements UserService
func (us *UserServiceImplement) GetAllData() ([]models.User, error) {
	users, errFindAll := us.userRepo.FindAll()
	if errFindAll != nil {
		return users, errFindAll
	}

	return users, nil
}

// GetAllData implements UserService
func (us *UserServiceImplement) GetDetailData(id string) ([]models.User, error) {

	users, errFindAll := us.userRepo.Find(id)
	if errFindAll != nil {
		return users, errFindAll
	}

	return users, nil
}

// New implements UserService
func (us *UserServiceImplement) New(userRequest request.UserRequest) (response.UserResponse, error) {
	var userResponse response.UserResponse
	user := models.User{
		Name: userRequest.Username,
	}

	errCreate := us.userRepo.Create(&user)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return userResponse, errCreate
	}

	userResponse.ID = user.ID
	userResponse.Name = user.Name

	return userResponse, nil
}

func (us *UserServiceImplement) Validation(userRequest request.UserRequest) error {
	var messageError string
	var isError bool

	// MANUAL VALIDATION
	if userRequest.Username == "" {
		messageError += "Username is required"
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}
