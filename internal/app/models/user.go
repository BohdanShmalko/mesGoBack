package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id          int
	Name        string
	Lastname    string
	MainPhoto   string
	Status      string
	AboutMe     string
	DefaultPath string
	Email       string
	Password    string
	Nickname    string
	RowPassword string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.RowPassword, validation.By(requiredIf(u.Password == "")), validation.Length(10, 40)),
		validation.Field(&u.Lastname, validation.Required),
		validation.Field(&u.Nickname, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		)
}

func (u *User) BeforeCreate() error {
	if len(u.RowPassword) == 0 {
		return errors.New("field RowPassword is empty")
	}
	ep, err := EncryptString(u.RowPassword)
	if err != nil {
		return err
	}
	u.Password = ep
	return nil
}

func EncryptString(row string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(row))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
