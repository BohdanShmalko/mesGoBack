package sqlstore_test

import (
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	"github.com/BohdanShmalko/mesGoBack/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, truncate := sqlstore.TestDB(t, databaseURL)
	defer truncate("users")
	s := sqlstore.New(db)

	user := models.TestUser(t)
	err := s.User().Create(user)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUserRepository_FindUser(t *testing.T) {
	db, truncate := sqlstore.TestDB(t, databaseURL)
	defer truncate("users")
	s := sqlstore.New(db)

	_, err := s.User().FindUser("badEmail", "badPassword")
	assert.EqualError(t, err, store.ErrRecordNotFind.Error())

	s.User().Create(models.TestUser(t))

	user, err := s.User().FindUser("email@gmail.com", "somepassword")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
