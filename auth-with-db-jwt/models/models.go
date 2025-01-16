package models

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Age      int                `json:"age"`
	Password string             `json:"password"`
}

func HashUserPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPassword(inputPassword string, databasePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(inputPassword))

	// when both of them are equal then we will return as true
	return err == nil
}

func CheckSigninRequiredFeilds(user User) error {
	if (user.Email == "" && user.Username == "") || user.Password == "" {
		return errors.New("required feilds not present")
	}

	return nil
}

func CheckSignupRequiredFeilds(user User) error {
	if user.Email == "" || user.Age == 0 || user.Password == "" || user.Username == "" {
		return errors.New("required feilds not present")
	}

	// now we will check if email is okay or not
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(user.Email) {
		return errors.New("malformed email")
	}

	if len(user.Password) <= 8 {
		return errors.New("password length less than 8")
	}

	if user.Age < 18 {
		return errors.New("user age less than 18")
	}

	return nil
}
