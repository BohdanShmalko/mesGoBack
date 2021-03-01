package models

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Name : "some name",
		Lastname : "some last name",
		Email : "email@gmail.com",
		RowPassword : "somepassword",
		DefaultPath : "/jjfjf/sjdj",
		Nickname : "nickname",
		AboutMe: "some about",
	}
}