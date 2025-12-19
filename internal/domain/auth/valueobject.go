package auth

import (
	"errors"
	"regexp"
	"strings"
)

type UserID string

func NewUserID(id string) (UserID, error) {
	if strings.TrimSpace(id) == "" {
		return "", errors.New("user id cannot be empty")
	}

	return UserID(id), nil
}

type Email string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (Email, error) {
	if value == "" {
		return "", errors.New("email cannot be empty")
	}

	if !emailRegex.MatchString(value) {
		return "", errors.New("invalid email format")
	}

	return Email(value), nil
}

type HashedPassword string

func NewHashedPassword(hash string) (HashedPassword, error) {
	if hash == "" {
		return "", errors.New("password hash cannot be empty")
	}
	return HashedPassword(hash), nil
}
