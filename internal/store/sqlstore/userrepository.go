package sqlstore

import (
	"database/sql"
	"github.com/BohdanShmalko/mesGoBack/helper"
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
)

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.BeforeCreate(); err != nil {
		return err
	}
	return ur.store.db.QueryRow(`INSERT INTO
Users (name, lastName, email, password, defaultPath, nickname, mainphoto, status, aboutme)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		user.Name, user.Lastname, user.Email, user.Password, user.DefaultPath, user.Nickname, user.MainPhoto, user.Status, user.AboutMe).Scan(
		&user.Id)

}

func (ur *UserRepository) FindUser(email, password string) (*models.User, error) {
	ep, err := models.EncryptString(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	var mainPhoto, status, aboutMe []byte
	if err := ur.store.db.QueryRow(`SELECT id, name, lastname, mainphoto, status, aboutme, defaultpath, email, password, nickname
FROM users WHERE email = $1 AND password = $2`, email, ep).Scan(
		&user.Id, &user.Name, &user.Lastname, &mainPhoto, &status, &aboutMe, &user.DefaultPath, &user.Email, &user.Password, &user.Nickname);
		err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFind
		}
		return nil, err
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}

func (ur *UserRepository) Find(id int) (*models.User, error) {
	user := &models.User{}
	var mainPhoto, status, aboutMe []byte
	if err := ur.store.db.QueryRow(`SELECT id, name, lastname, mainphoto, status, aboutme, defaultpath, email, password, nickname
FROM users WHERE id = $1`, id).Scan(
		&user.Id, &user.Name, &user.Lastname, &mainPhoto, &status, &aboutMe, &user.DefaultPath, &user.Email, &user.Password, &user.Nickname);
		err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFind
		}
		return nil, err
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}
