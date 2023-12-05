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

	//CONTROLLA CHE IL COMMENTO E' DELL'UTENTE E CHE NON COMMENTA FOTO DI CHI LO HA BANNATO

	comment.PhotoID = pid
	comment.UserID = uid
	comment.Date = utils.NowFormat()

	res, err := rt.db.CommentPhoto(comment)
	if res == database.NO_ROWS {
		http.Error(w, utils.ErrPhotoNotFound.Error()+" or "+utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
