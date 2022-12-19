package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mindmap-go/app/models"
	"mindmap-go/internal/database"
	"mindmap-go/utils"
)

type UserRepo struct {
	DB *database.Database
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAll() ([]*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByAccount(account *models.Account) (*models.User, error)
	UpdateUser(user *models.User, req *models.UserUpdate) error
	DeleteUser(user *models.User) error
}

func NewUserRepository(database *database.Database) UserRepository {
	return &UserRepo{
		DB: database,
	}
}

func (u UserRepo) CreateUser(user *models.User) error {
	acc := models.Account{Email: user.Account.Email}
	err := u.DB.Connection.Find(&acc).Or("username = ?", user.Account.Username).Error
	if err == nil {
		return errors.Wrap(utils.DuplicateEntryError, "the user with such credentials already exists")
	}
	return u.DB.Connection.Create(&user).Error
}

func (u UserRepo) GetAll() ([]*models.User, error) {
	var res []*models.User
	err := u.DB.Connection.Joins("Account").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u UserRepo) GetUserByID(id int) (*models.User, error) {
	var res *models.User
	err := u.DB.Connection.Joins("Account").First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u UserRepo) GetUserByAccount(account *models.Account) (*models.User, error) {
	var res *models.User
	err := u.DB.Connection.Find(&res).Where("account_id = ?", account.ID).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

func (u UserRepo) UpdateUser(user *models.User, req *models.UserUpdate) error {
	return u.DB.Connection.Save(&user).Error
}

func (u UserRepo) DeleteUser(user *models.User) error {
	return u.DB.Connection.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&user).Error
		if err != nil {
			return err
		}
		return tx.Delete(&user.Account).Error
	})
}
