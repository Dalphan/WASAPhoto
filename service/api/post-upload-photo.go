package api

import (
	"encoding/json"
	"fmt"
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

	// Non si puo salvare cosi l'immagine, la risposta è completamente occupata dall'immagine in base64
	photo.Image = utils.BytesToBase64(imageBytes)
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
	photo.PhotoID = uint(PID)

	fmt.Println("PHOTO ID", photo.PhotoID)

	utils.SetHeaderJson(w)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
