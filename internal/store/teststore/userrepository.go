package teststore

import (
	"github.com/BohdanShmalko/mesGoBack/helper"
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
)

type UserRepository struct {
	store *Store
	users map[int]*models.User
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

	id := len(ur.users) + 1
	ur.users[id] = user
	user.Id = id

	return nil
}

func (ur *UserRepository) FindUser(email, password string) (*models.User, error) {
	ep, err := models.EncryptString(password)
	if err != nil {
		return nil, err
	}

	var mainPhoto, status, aboutMe []byte

	user := &models.User{}
	ok := false
	for _, usr := range ur.users {
		if usr.Email == email && usr.Password == ep {
			user = usr
			ok = true
			break
		}
	}

	if !ok {
		return nil, store.ErrRecordNotFind
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}

func (ur *UserRepository) Find(id int) (*models.User, error) {
	var mainPhoto, status, aboutMe []byte

	user := &models.User{}
	ok := false
	for _, usr := range ur.users {
		if usr.Id == id {
			user = usr
			ok = true
			break
		}
	}

	if !ok {
		return nil, store.ErrRecordNotFind
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}