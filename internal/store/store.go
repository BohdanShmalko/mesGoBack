package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Store struct {
	config *Config
	db *sql.DB
	userRepository *UserRepository
}

func New(sConf *Config) *Store {
	return &Store{
		config: sConf,
	}
}

func (st *Store) Open() error{
	log.Println(st.config.DatabaseUrl)
	db, err := sql.Open("postgres", st.config.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	st.db = db
	return nil
}

func (st *Store) Close() {
	st.db.Close()
}

func (st *Store) User() *UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}
	st.userRepository = &UserRepository{store: st}
	return st.userRepository
}