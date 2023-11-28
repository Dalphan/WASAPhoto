package api

import (
	"net/http"
	"strconv"

	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := strconv.Atoi(ps.ByName("id"))
	if err != nil { //Error getting the user ID
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bid, err := strconv.Atoi(ps.ByName("bid"))
	if err != nil { //Error getting the user ID to ban
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if uid == bid {
		http.Error(w, "you can't unban yourself", http.StatusBadRequest)
		return
	}

}
