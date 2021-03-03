package teststore

import (
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{
	}
}

func (st *Store) User() store.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{
		store: st,
		users: make(map[int]*models.User),
	}
	return st.userRepository
}