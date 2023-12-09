package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	pid, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	cid, err := utils.GetHttpParam(w, ps, "cid")
	if err != nil {
		return
	}

	uid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	// The user can only delete their own comments
	c, res, err := rt.db.GetCommentById(cid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, utils.ErrCommentNotFound.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case database.SUCCESS:
		if c.UserID != uid { // The comment is from another user
			http.Error(w, utils.ErrUnauthoraized.Error(), http.StatusUnauthorized)
			return
		}
	}

	res, err = rt.db.UncommentPhoto(cid, pid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrPhotoNotFound.Error(), http.StatusNotFound)
		return
	}

	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode("Comment deleted succesfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
