package services

import (
	"errors"
	"fmt"

	"github.com/Kvothe838/drivers-api/db"
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

	user, err := db.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error getting user by username: %v", err)
	}

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
func encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func compareHashAndPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func SignUp(username, password string) (*model.User, error) {
	alreadyExistingUser, err := db.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error getting user by username: %v", err)
	}

	if alreadyExistingUser != nil {
		return nil, UserAlreadyExists
	}

	hashedPassword, err := encrypt(password)
	if err != nil {
		return nil, fmt.Errorf("error encrypting password: %v", err)
	}

	newUser := model.User{
		Username: username,
		Hash:     hashedPassword,
	}

	savedUser, err := db.SaveUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("error saving user in db: %v", err)
	}

	return savedUser, nil
}

func UserHasPermission(userId int64, permissionName string) (*bool, error) {
	return db.UserHasPermission(userId, permissionName)
}
