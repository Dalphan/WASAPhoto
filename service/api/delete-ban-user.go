package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil { //Error getting the user ID
		return
	}

	bid, err := utils.GetHttpParam(w, ps, "bid")
	if err != nil { //Error getting the user ID to unban
		return
	}

	if uid == bid {
		http.Error(w, "you can't unban yourself", http.StatusBadRequest)
		return
	}

	res, err := rt.db.UnbanUser(uid, bid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}

	user, res, err := rt.db.FindUserByID(bid)
	// if res == database.NO_ROWS {
	// 	http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
	// 	return
	// }
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unbanned user returned succesfully
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
