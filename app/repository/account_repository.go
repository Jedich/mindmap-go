package repository

import (
	"mindmap-go/app/models"
	"mindmap-go/internal/database"
)

type AccountRepo struct {
	DB *database.Database
}

type AccountRepository interface {
	CreateAccount(account *models.Account) error
	GetAll() ([]*models.Account, error)
	GetAccountByID(id int) (*models.Account, error)
	UpdateAccount(account *models.Account, req *models.AccountUpdate) error
	DeleteAccount(account *models.Account) error
}

func NewAccountRepository(database *database.Database) AccountRepository {
	return &AccountRepo{
		DB: database,
	}
}

func (a AccountRepo) CreateAccount(account *models.Account) error {
	return a.DB.Connection.Create(&account).Error
}

func (a AccountRepo) GetAll() ([]*models.Account, error) {
	var res []*models.Account
	err := a.DB.Connection.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AccountRepo) GetAccountByID(id int) (*models.Account, error) {
	var res *models.Account
	err := a.DB.Connection.First(&res, id).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AccountRepo) UpdateAccount(account *models.Account, req *models.AccountUpdate) error {
	return a.DB.Connection.Model(&account).
		Updates(models.Account{Username: req.Username, PasswordHash: req.PasswordHash}).Error
}

func (a AccountRepo) DeleteAccount(account *models.Account) error {
	return a.DB.Connection.Delete(&account).Error
}
