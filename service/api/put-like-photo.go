package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	// Check if the photo exists
	photo, res, err := rt.db.GetPhotoById(pid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, utils.ErrPhotoNotFound.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Photo returned succesfully
	// Check if photo owner banned the one posting the comment
	if CheckBanned(w, rt, photo.UserID, lid, utils.ErrPhotoNotFound) {
		return
	}

	like, res, err := rt.db.LikePhoto(pid, lid)
	switch res {
	case database.UNIQUE_FAILED:
		w.WriteHeader(http.StatusOK)
		utils.SetHeaderJson(w)
		like.PhotoID = pid
		like.UserID = lid
		err = json.NewEncoder(w).Encode(like)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	case database.SUCCESS:
		w.WriteHeader(http.StatusCreated)
		utils.SetHeaderJson(w)
		err = json.NewEncoder(w).Encode(like)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	case database.NO_ROWS:
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
