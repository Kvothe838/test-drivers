package services

import (
	"errors"
	"fmt"

	"github.com/Kvothe838/drivers-api/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserOrPasswordIncorrect = errors.New("Username or password incorrect")
	UserAlreadyExists       = errors.New("Username already exists")
)

var allUsers []model.User = make([]model.User, 0)

func Login(username, password string) (*model.User, error) {
	hashedPassword, err := encrypt(password)
	if err != nil {
		fmt.Printf("error encrypting password: %v\n", err)
		return nil, err
	}
	user := getUserByUsername(username)

	if user == nil {
		return nil, UserOrPasswordIncorrect
	}

	isPasswordCorrect := compareHashAndPassword(hashedPassword, user.Hash)

	if isPasswordCorrect {
		return user, nil
	} else {
		return nil, UserOrPasswordIncorrect
	}
}

func getUserByUsername(username string) *model.User {
	for _, user := range allUsers {
		if user.Username == username {
			return &user
		}
	}

	return nil
}

func encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func compareHashAndPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func SignUp(username, password string) (*model.User, error) {
	alreadyExistingUser := getUserByUsername(username)
	if alreadyExistingUser != nil {
		return nil, UserAlreadyExists
	}

	hashedPassword, err := encrypt(password)
	if err != nil {
		fmt.Printf("error encrypting password: %v", err)
		return nil, err
	}

	newUser := model.User{
		Username: username,
		Hash:     hashedPassword,
	}

	allUsers = append(allUsers, newUser)

	return &newUser, nil
}
