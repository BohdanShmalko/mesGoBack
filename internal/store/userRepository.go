package store

import "github.com/BohdanShmalko/mesGoBack/internal/app/models"

type UserRepository interface {
	Create(*models.User) error
	FindUser(string, string) (*models.User, error)
}
