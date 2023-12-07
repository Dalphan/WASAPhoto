package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	pid, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	uid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}
	uid += 1

	// Check if photo exists
	_, p_res, err := rt.db.GetPhotoById(pid)
	switch p_res {
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case database.NO_ROWS:
		http.Error(w, utils.ErrPhotoNotFound.Error(), http.StatusNotFound)
		return
	}

	comments, res, err := rt.db.GetCommentsByPhoto(pid)
	switch res {
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case database.NO_ROWS:
		http.Error(w, utils.ErrCommentNotFound.Error(), http.StatusNotFound)
		return
	case database.SUCCESS:
		utils.SetHeaderJson(w)
		err := json.NewEncoder(w).Encode(comments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
