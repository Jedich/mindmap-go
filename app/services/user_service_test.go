package services

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
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
