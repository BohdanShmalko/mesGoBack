package sqlstore

import (
	"database/sql"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	_ "github.com/lib/pq"
)

//TODO in main database varchar(255)
type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db : db,
	}
}

func (st *Store) User() store.UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{store: st}
	return st.userRepository
}