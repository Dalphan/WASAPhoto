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

	_, _ = rt.db.UnbanUser(uid, bid)
	// Controllare risposta

	user, res, err := rt.db.FindUserByID(bid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unbanned user returned succesfully
	utils.SetHeaderJson(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
