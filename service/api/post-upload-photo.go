package api

import (
	"encoding/json"
	"io"
	"net/http"

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
	photo.UserID = uid

	photo.Caption = r.URL.Query().Get("caption")
	if len(photo.Caption) > 1000 {
		http.Error(w, utils.ErrCaptionTooBig.Error(), http.StatusNotAcceptable)
		return
	}

	imageBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(imageBytes) > 20971520 {
		http.Error(w, utils.ErrPhotoTooBig.Error(), http.StatusNotAcceptable)
		return
	}

	// Non si puo salvare cosi l'immagine, la risposta è completamente occupata dall'immagine in base64
	photo.Image = utils.BytesToBase64(imageBytes)
	photo.Timestamp = utils.NowFormat()

	PID, res, err := rt.db.CreatePhoto(photo)
	if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	photo.PhotoID = PID

	utils.SetHeaderJson(w)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}