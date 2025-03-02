package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/nethsaraPrabash/chat-app-go/src/models"
	"github.com/nethsaraPrabash/chat-app-go/src/repository"
)


func isUserExists(email, username string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err == nil && user != nil {
		return user, errors.New("email has already been taken")
	}

	user, err = repository.GetUserByUsername(username)
	if err == nil && user != nil {
		return user, errors.New("username has already been taken")
	}

	return nil, nil
}


func RegisterUser(user *models.User) error {
	if _, err := isUserExists(user.Email, user.Username); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return repository.Register(user)
}

func Login(user *models.User) error {
	existingUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return errors.New("invalid credentials")
	}

	user.ID = existingUser.ID
	return nil
}
