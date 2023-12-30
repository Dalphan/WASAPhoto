package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil {
		return
	}

	authid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	page_str := r.URL.Query().Get("page")
	if page_str == "" {
		page_str = "0"
	}

	page, err := strconv.Atoi(page_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the User has banned the one searching
	if CheckBanned(w, rt, authid, uid, utils.ErrUserNotFound) {
		return
	}

	photos, res, err := rt.db.GetUserPhotos(uid, page)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
