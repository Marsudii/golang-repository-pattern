package repositories

import "go-fiber/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	Find(id string) ([]models.User, error)
	Create(user *models.User) error
}
