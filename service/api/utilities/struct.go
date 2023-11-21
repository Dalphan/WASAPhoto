package utilities

import "errors"

// `

type User struct {
	UserID         uint   `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	FollowersCount int    `json:"followersCount"`
	FollowingCount int    `json:"followingCount"`
	PhotoCount     int    `json:"photoCount"`
}

type Photo struct {
	PhotoID      uint   `json:"id"`
	UserID       uint   `json:"user"`
	Image        []byte `json:"image"`
	Timestamp    string `json:"timestamp"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

type Comment struct {
	CommentID uint   `json:"id"`
	Text      string `json:"text"`
	UserID    uint   `json:"user"`
	PhotoID   uint   `json:"photo"`
	Date      string `json:"date"`
}

var ErrUsernameMissing = errors.New("missing required username")
var ErrUsernameNotValid = errors.New("the username doesnt match the required pattern")

func validateUsername(username string) error {
	if len(username) == 0 {
		return ErrUsernameMissing
	}
	if len(username) < 3 || len(username) > 16 {
		return ErrUsernameNotValid
	}
	return nil
}
