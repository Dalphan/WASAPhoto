package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	pid, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	lid, err := utils.GetHttpParam(w, ps, "lid")
	if err != nil {
		return
	}

	_, err = utils.GetAuthorization(w, r, lid)
	if err != nil {
		return
	}

	res, err := rt.db.UnlikePhoto(pid, lid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error()+" or "+utils.ErrPhotoNotFound.Error(), http.StatusNotFound)
		return
	}

	var like utils.Like
	like.PhotoID = pid
	like.UserID = lid
	utils.SetHeaderText(w)
	err = json.NewEncoder(w).Encode(like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
