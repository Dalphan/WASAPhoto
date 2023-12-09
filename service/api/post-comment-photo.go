package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	pid, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	var comment utils.Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !comment.Validate() {
		http.Error(w, utils.ErrCommentNotValid.Error(), http.StatusNotAcceptable)
		return
	}

	uid, err := utils.GetAuthorization(w, r)
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
	if CheckBanned(w, rt, photo.UserID, uid, utils.ErrPhotoNotFound) {
		return
	}

	comment.PhotoID = pid
	comment.UserID = uid
	comment.Date = utils.NowFormat()

	cid, res, err := rt.db.CommentPhoto(comment)
	if res == database.NO_ROWS {
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentID = cid
	w.WriteHeader(http.StatusCreated)
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
