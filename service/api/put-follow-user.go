package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil {
		return
	}

	// Check for Authorization
	if _, err = utils.GetAuthorization(w, r, uid); err != nil {
		return
	}

	fid, err := utils.GetHttpParam(w, ps, "fid")
	if err != nil {
		return
	}

	if uid == fid {
		http.Error(w, "you can't follow yourself", http.StatusBadRequest)
		return
	}

	switch res, err := rt.db.FollowUser(uid, fid); res {
	case database.UNIQUE_FAILED:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "the given user was already followed\n")
		return
	case database.NO_ROWS:
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The ban operation succeded, return the followed user

	user, res, err := rt.db.FindUserByID(fid)
	// if res == database.NO_ROWS { // Technically if the follow is okay, the user exists and this is not needed
	// 	http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
	// 	return
	// }
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Banned user returned succesfully
	utils.SetHeaderJson(w)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
