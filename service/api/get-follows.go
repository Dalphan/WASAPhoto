package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

type db_follow func(int) ([]utils.User, int, error)

func getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params, rt *_router, follow db_follow, mes string) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil {
		return
	}

	// Check for Authorization
	auth, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	// Check if uid banned auth
	if CheckBanned(w, rt, uid, auth, utils.ErrUserNotFound) {
		return
	}

	users, res, err := follow(uid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		http.Error(w, mes+" or user not found", http.StatusNotFound)
		return
	}

	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}

func (rt *_router) getFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getFollows(w, r, ps, rt, rt.db.GetFollowings, "followings")
}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	getFollows(w, r, ps, rt, rt.db.GetFollowers, "followers")
}
