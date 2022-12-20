package services

import (
	"crypto/sha256"
	"fmt"
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
)

type UserSvc struct {
	Repo        repository.UserRepository
	accountRepo repository.AccountRepository
}

type UserService interface {
	Register(form *RegisterForm) (*models.User, error)
	Deregister(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByAccount(account *models.Account) (*models.User, error)
	UpdateUser(user *models.User, req *models.UserUpdate) error
	AuthorizeUser(l *LoginForm) (*models.User, error)
	Hash(text string) string
}

func NewUserService(repo repository.UserRepository, acc repository.AccountRepository) UserService {
	return &UserSvc{
		Repo:        repo,
		accountRepo: acc,
	}
}

func (u *UserSvc) Register(form *RegisterForm) (*models.User, error) {
	user := models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Account: models.Account{
			Username:     form.Username,
			Email:        form.Email,
			PasswordHash: u.Hash(form.Password),
		},
	}
	if err := u.Repo.CreateUser(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserSvc) Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *UserSvc) GetAllUsers() ([]*models.User, error) {
	return u.Repo.GetAll()
}

func (u *UserSvc) AuthorizeUser(l *LoginForm) (*models.User, error) {
	return u.Repo.GetUserByCredentials(&models.Account{
		Email:        l.Email,
		PasswordHash: l.Password,
	})
}

func (u *UserSvc) Deregister(user *models.User) error {
	return u.Repo.DeleteUser(user)
}

func (u *UserSvc) GetUserByID(id int) (*models.User, error) {
	return u.Repo.GetUserByID(id)
}

func (u *UserSvc) GetUserByAccount(account *models.Account) (*models.User, error) {
	return u.Repo.GetUserByAccount(account)
}

func (u *UserSvc) UpdateUser(user *models.User, req *models.UserUpdate) error {
	return u.Repo.UpdateUser(user, req)
}
