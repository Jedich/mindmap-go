package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	GetUserByCredentials(account *models.Account, password string) (*models.User, error)
	UpdateUser(user *models.User, req *models.UserUpdate) error
	DeleteUser(user *models.User) error
}

func NewUserRepository(database *database.Database) UserRepository {
	return &UserRepo{
		DB: database,
	}
}

func (u *UserRepo) CreateUser(user *models.User) error {
	if err := u.hasUserByCredentials(&user.Account); err != nil {
		return err
	}
	return u.DB.Connection.Create(&user).Error
}

func (u *UserRepo) hasUserByCredentials(account *models.Account) error {
	var acc *models.Account
	if err := u.DB.Connection.Where("email = ? OR username = ?", account.Email, account.Username).First(&acc).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return &utils.DuplicateEntryError{Message: "User with such credentials already exists."}
}

func (u *UserRepo) GetUserByCredentials(account *models.Account, password string) (*models.User, error) {
	var acc *models.Account
	err := u.DB.Connection.Where("email = ?", account.Email).First(&acc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &utils.UnauthorizedEntryError{Message: "Your login credentials are invalid."}
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(acc.PasswordHash, []byte(password)); err != nil {
		return nil, &utils.UnauthorizedEntryError{Message: "Your login credentials are invalid. (password)"}
	}
	return u.GetUserByAccount(acc)
}

func (u *UserRepo) GetAll() ([]*models.User, error) {
	var res []*models.User
	err := u.DB.Connection.Joins("Account").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserRepo) GetUserByID(id int) (*models.User, error) {
	var res *models.User
	err := u.DB.Connection.Joins("Account").First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserRepo) GetUserByAccount(account *models.Account) (*models.User, error) {
	var res models.User
	err := u.DB.Connection.Find(&res).Where("account_id = ?", account.ID).Error
	res.ID = account.ID
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (u *UserRepo) UpdateUser(user *models.User, req *models.UserUpdate) error {
	return u.DB.Connection.Save(&user).Error
}

func (u *UserRepo) DeleteUser(user *models.User) error {
	return u.DB.Connection.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&user).Error
		if err != nil {
			return err
		}
		return tx.Delete(&user.Account).Error
	})
}
