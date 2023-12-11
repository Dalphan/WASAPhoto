package utils

import (
	"encoding/base64"
	"strings"
)

// `

type User struct {
	UserID         int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	FollowersCount int    `json:"followersCount"`
	FollowingCount int    `json:"followingCount"`
	PhotoCount     int    `json:"photoCount"`
}

type Photo struct {
	PhotoID      int    `json:"id"`
	UserID       int    `json:"user"`
	Image        string `json:"image"`
	Timestamp    string `json:"timestamp"`
	Caption      string `json:"caption"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

type Comment struct {
	CommentID int    `json:"id"`
	Text      string `json:"text"`
	UserID    int    `json:"user"`
	PhotoID   int    `json:"photo"`
	Date      string `json:"date"`
}

type Like struct {
	UserID  int `json:"user"`
	PhotoID int `json:"photo"`
}

func (c Comment) Validate() bool {
	c.Text = strings.TrimSpace(c.Text) // Removes spaces at the beginning and the end. If the string is empty, the length becomes 0
	return len(c.Text) > 1 && len(c.Text) < 5000
}

func BytesToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// FromBase64 decodes the image data from Base64.
// func (p *Photo) FromBase64() error {
// 	imageBytes, err := base64.StdEncoding.DecodeString(p.Image)
// 	if err != nil {
// 		return err
// 	}
// 	p.Image = string(imageBytes)
// 	return nil
// }
