package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfileById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil {
		return
	}

	auth, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	// If the user making the request is banned by the searched user, the response will not be found
	if CheckBanned(w, rt, uid, auth, utils.ErrUserNotFound) {
		return
	}

	user, res, err := rt.db.FindUserByID(uid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case database.SUCCESS:
		utils.SetHeaderJson(w)
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
