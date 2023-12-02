package utils

import "encoding/base64"

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
	Image        string `json:"image"`
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
