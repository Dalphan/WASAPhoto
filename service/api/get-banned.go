package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil {
		return
	}

	// Check for Authorization, the list is visibly only to the author
	_, err = utils.GetAuthorization(w, r, uid)
	if err != nil {
		return
	}

	banned, res, err := rt.db.GetBanned(uid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if len(banned) == 0 {
	// 	http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
	// 	return
	// }

	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(banned)
	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
