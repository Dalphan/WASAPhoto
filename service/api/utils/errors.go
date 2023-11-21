package utils

import "errors"

var ErrUsernameMissing = errors.New("missing required username")
var ErrUsernameNotValid = errors.New("the username doesnt match the required pattern")
