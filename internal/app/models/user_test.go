package models_test

import (
	"github.com/BohdanShmalko/mesGoBack/internal/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	userCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "valid user",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "not valid email",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.Email = "some bad email"
				return tu
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.RowPassword = ""
				return tu
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.RowPassword = "lol"
				return tu
			},
			isValid: false,
		},
		{
			name: "long password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.RowPassword = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
				return tu
			},
			isValid: false,
		},
		{
			name: "with encrypt password",
			u: func() *models.User {
				tu := models.TestUser(t)
				tu.RowPassword = ""
				tu.Password = "$2a$04$hErb9LF8Fkn7wfI3zAQLdeYExLaHhpc4WVhr2ymo66bmC57fQXI1y"
				return tu
			},
			isValid: true,
		},
	}

	for _, tc := range userCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.Validate())
}
