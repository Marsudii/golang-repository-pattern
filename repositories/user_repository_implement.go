package repositories

import (
	"go-fiber/database"
	"go-fiber/models"
	"strconv"
)

type UserRepositoryImplement struct {
}

// Create implements UserRepository
func (*UserRepositoryImplement) Create(user *models.User) error {
	db := database.DB.Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

// FindAll implements UserRepository
func (*UserRepositoryImplement) FindAll() ([]models.User, error) {
	//AMBIL MODEL USER
	var users []models.User
	// ORM KE DB ALL DATA
	db := database.DB.Find(&users)
	//CHECKING
	if db.Error != nil {
		return users, db.Error
	}

	return users, nil
}

// DETAIL DATA
func (*UserRepositoryImplement) Find(id string) ([]models.User, error) {

	// PARAMS CONVERT TO INTEGER
	convertIdtoInteger, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var users []models.User
	// ORM KE DB ALL DATA
	db := database.DB.Find(&users, convertIdtoInteger)
	if db.Error != nil {
		return users, db.Error
	}
	//
	return users, nil
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplement{}
}
