package teststore_test

import (
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/BohdanShmalko/mesGoBack/internal/store"
	"github.com/BohdanShmalko/mesGoBack/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	user := models.TestUser(t)
	err := s.User().Create(user)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUserRepository_FindUser(t *testing.T) {
	s := teststore.New()

	_, err := s.User().FindUser("badEmail", "badPassword")
	assert.EqualError(t, err, store.ErrRecordNotFind.Error())

	s.User().Create(models.TestUser(t))

	user, err := s.User().FindUser("email@gmail.com", "somepassword")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	_, err := s.User().Find(-1)
	assert.EqualError(t, err, store.ErrRecordNotFind.Error())

	testUser := models.TestUser(t)
	s.User().Create(testUser)

	user, err := s.User().Find(testUser.Id)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
