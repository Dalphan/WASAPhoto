package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		http.Error(w, "you can't unfollow yourself", http.StatusBadRequest)
		return
	}

	res, err := rt.db.UnfollowUser(uid, fid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}

	user, res, err := rt.db.FindUserByID(fid)
	// if res == database.NO_ROWS {
	// 	http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
	// 	return
	// }
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unfollowed user returned succesfully
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
