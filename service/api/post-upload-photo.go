package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	var photo utils.Photo
	photo.UserID = uint(uid)

	imageBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(imageBytes) > 20971520 {
		http.Error(w, utils.ErrPhotoTooBig.Error(), http.StatusNotAcceptable)
		return
	}

	photo.Image = imageBytes
	photo.Timestamp = time.Now().Format("2006-01-02T15:04:05Z")

	PID, res, err := rt.db.CreatePhoto(photo)
	if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SetHeaderJson(w)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(PID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
