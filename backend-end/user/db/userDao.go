package db

import (
	"github.com/haitao-sun03/donation/backend-end/user/config"
	"github.com/haitao-sun03/donation/backend-end/user/model"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	db := config.GetDB()
	return &UserDao{db}
}

func (userDao *UserDao) GetUserByAddress(address string) (model.UserModel, error) {
	var user model.UserModel
	if err := userDao.db.Where("address = ?", address).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (userDao *UserDao) CreateUser(user model.UserModel) error {
	return userDao.db.Create(&user).Error
}
