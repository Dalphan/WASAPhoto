package api

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
