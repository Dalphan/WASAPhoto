package utils

import "errors"

var ErrUsernameMissing = errors.New("missing required username")
var ErrUsernameNotValid = errors.New("the username doesnt match the required pattern")
var ErrUserNotFound = errors.New("user not found")
var ErrUnauthorazied = errors.New("unauthorazied operation")

var ErrUsernameTaken = errors.New("the username is already taken")
