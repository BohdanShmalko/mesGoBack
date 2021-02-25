package store

import (
	"github.com/BohdanShmalko/mesGoBack/helper"
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := ur.store.db.QueryRow(`INSERT INTO
Users (name, lastName, email, password, defaultPath, nickname, mainphoto, status, aboutme)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		user.Name, user.Lastname, user.Email, user.Password, user.DefaultPath, user.Nickname, user.MainPhoto, user.Status, user.AboutMe).Scan(
		&user.Id); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindUser(email, password string) (*models.User, error) {
	user := &models.User{}
	var mainPhoto, status, aboutMe []byte
	if err := ur.store.db.QueryRow(`SELECT id, name, lastname, mainphoto, status, aboutme, defaultpath, email, password, nickname
FROM users WHERE email = $1 AND password = $2`, email, password).Scan(
	&user.Id, &user.Name, &user.Lastname, &mainPhoto, &status, &aboutMe, &user.DefaultPath, &user.Email, &user.Password, &user.Nickname);
	err != nil {
		return nil, err
	}

	user.MainPhoto = helper.Get(mainPhoto)
	user.Status = helper.Get(status)
	user.AboutMe = helper.Get(aboutMe)
	return user, nil
}
