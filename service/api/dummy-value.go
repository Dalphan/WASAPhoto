package api

import (
	"github.com/Dalphan/WASAPhoto/service/utils"
)

var users = []utils.User{
	{
		UserID:         1,
		Username:       "Babba",
		Name:           "Babbino",
		Surname:        "Babbetto",
		FollowersCount: 13,
		FollowingCount: 137,
		PhotoCount:     2,
	},
	{
		UserID:         2,
		Username:       "Trambusto",
		Name:           "Tommaso",
		Surname:        "Brusto",
		FollowersCount: 193,
		FollowingCount: 27,
		PhotoCount:     92,
	},
}
