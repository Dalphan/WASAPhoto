package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)
	var uid int

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil { // Error getting the user ID
		return
	}

	// Check for Authorization
	if _, err = utils.GetAuthorization(w, r, uid); err != nil {
		return
	}

	bid, err := utils.GetHttpParam(w, ps, "bid")
	if err != nil { // Error getting the user ID to unban
		return
	}
	if uid == bid {
		http.Error(w, "you can't ban yourself", http.StatusBadRequest)
		return
	}

	switch res, err := rt.db.BanUser(uid, bid); res {
	case database.UNIQUE_FAILED:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "the given user was already banned\n")
		return
	case database.NO_ROWS:
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The ban operation succeded, return the banned user

	user, res, err := rt.db.FindUserByID(bid)
	// if res == database.NO_ROWS { // Technically if the ban is okay, the user exists and this is not needed
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
