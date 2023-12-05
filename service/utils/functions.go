package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NowFormat() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

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

func GetHttpParam(w http.ResponseWriter, ps httprouter.Params, name string) (int, error) {
	uid, err := strconv.Atoi(ps.ByName(name))
	if err != nil { // Error getting the param from the path
		http.Error(w, err.Error(), http.StatusBadRequest)
		return uid, err
	}
	return uid, nil
}

func SetHeaderText(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
}

func SetHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAuthorization(w http.ResponseWriter, r *http.Request, id ...int) (int, error) {
	uid, err := strconv.Atoi(r.Header.Get("Authorization"))

	if err != nil || (len(id) > 0 && uid != id[0]) {
		http.Error(w, ErrUnauthorazied.Error(), http.StatusUnauthorized)
		return 0, ErrUnauthorazied
	}

	return uid, nil
}
