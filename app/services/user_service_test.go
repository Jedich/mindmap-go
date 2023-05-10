package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"mindmap-go/app/repository"
	"testing"
	"time"
)

func TestHashing(t *testing.T) {
	str := "some string"
	svc := UserSvc{}
	x1, err := svc.Hash(str)
	if err != nil {
		t.Error(err)
	}

	x2, err := svc.Hash("some string")
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, bcrypt.CompareHashAndPassword(x1, []byte("some string")))
	assert.Nil(t, bcrypt.CompareHashAndPassword(x2, []byte("some string")))
	assert.NotEqual(t, x1, x2)
}

func TestGetAccountByIDExists(t *testing.T) {
	db, mock := createMockDB(t)
	rows := sqlmock.NewRows([]string{`created_at`, `updated_at`, `deleted_at`, `name`, `desc`, `creator_id`}).
		AddRow(time.Now(), time.Now(), nil, "Unnamed map", "", 0)
	mock.ExpectQuery("SELECT(.*)").WithArgs(0).
		WillReturnRows(rows)

	l, err := NewUserService(repository.NewUserRepository(db), repository.NewAccountRepository(db)).GetUserByID(0)
	assert.NoError(t, err)
	assert.NotEmpty(t, l)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
