package teststore

import (
	"github.com/BohdanShmalko/mesGoBack/helper"
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
)

type UserRepository struct {
	store *Store
	users map[AuthKey]*models.User
}

type AuthKey struct {
	email string
	password string
}

func (ur *UserRepository) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	ur.users[AuthKey{
		email: user.Email,
		password: user.Password,
	}] = user
	user.Id = len(ur.users)

	return nil
}

func (ur *UserRepository) FindUser(email, password string) (*models.User, error) {
	ep, err := models.EncryptString(password)
	if err != nil {
		return nil, err
	}

	var mainPhoto, status, aboutMe []byte
	user, ok := ur.users[AuthKey{
		email: email,
		password: ep,
	}]
	if !ok {
		return nil, store.ErrRecordNotFind
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}
