package repository

import (
	"clockwork-server/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id uint64) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindAll() ([]model.User, error)
	Update(user model.User) (model.User, error)
	Delete(userID uint64) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &repository{db}
}

func (r *repository) Save(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(id uint64) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]model.User, error) {
	var user []model.User
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(userID uint64) (string, error) {
	err := r.db.Delete(&model.User{}, userID).Error
	if err != nil {
		return "error delete user", err
	}

	return "success delete user", nil
}
