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

	// Accept alphanumeric username between 3 and 16 characters, with . and _
	match, err := regexp.MatchString("^[a-zA-Z0-9._]{3,16}$", username)
	if err != nil || !match { // The username doesn't match the required pattern
		return ErrUsernameNotValid
	} else { // The username is valid
		return nil
	}
}

// Write the right response code based on the validity of the username
func HttpValidateUsername(w http.ResponseWriter, username string) bool {
	err := ValidateUsername(username)
	if errors.Is(err, ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return false
	}
	if errors.Is(err, ErrUsernameMissing) || err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}
