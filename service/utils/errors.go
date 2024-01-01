package utils

import "errors"

var ErrUsernameMissing = errors.New("missing required username")
var ErrUsernameNotValid = errors.New("the username doesnt match the required pattern")
var ErrUserDetailsNotValid = errors.New("the user's details dont match the required pattern")
var ErrUserNotFound = errors.New("user not found")
var ErrUnauthoraized = errors.New("unauthorazied operation")
var ErrUsernameTaken = errors.New("the username is already taken")

var ErrPhotoTooBig = errors.New("the provided photo exceeds the maximum dimension")
var ErrPhotoNotFound = errors.New("photo not found")
var ErrNotPhoto = errors.New("the given file is not an image")
var ErrCaptionTooBig = errors.New("the caption exceeds the maximun length")

var ErrCommentNotValid = errors.New("the given comment doesn't match the required length")
var ErrCommentNotFound = errors.New("comment not found")

var ErrEncodingJson = errors.New("error encoding json")
