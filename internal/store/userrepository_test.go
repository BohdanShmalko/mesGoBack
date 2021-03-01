package store_test

import (
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, truncate := store.TestStore(t, databaseURL)
	defer truncate("users")

	user, err := s.User().Create(models.TestUser(t))
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUserRepository_FindUser(t *testing.T) {
	s, truncate := store.TestStore(t, databaseURL)
	defer truncate("users")

	_, err := s.User().FindUser("badEmail", "badPassword")
	assert.Error(t, err)

	s.User().Create(models.TestUser(t))

	user, err := s.User().FindUser("email@gmail.com", "somepassword")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
