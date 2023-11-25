package utils

import (
	"errors"
	"net/http"
	"regexp"
)

// Database checks if username is between 3 and 16 characters
func ValidateUsername(username string) error {
	if len(username) == 0 {
		return ErrUsernameMissing
	}
	// if len(username) < 3 || len(username) > 16 {
	// 	return ErrUsernameNotValid
	// }

	match, err := regexp.MatchString("^[a-zA-Z0-9._]{3,16}$", username)
	if err != nil {
		return err
	}
	if match {
		return nil
	}
	return ErrUsernameNotValid
}

func HttpValidateUsername(w http.ResponseWriter, username string) bool {
	err := ValidateUsername(username)
	if errors.Is(err, ErrUsernameMissing) || err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	} else if errors.Is(err, ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return false
	}

	return true
}
