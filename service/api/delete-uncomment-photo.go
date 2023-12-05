package api

import (
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	_, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	cid, err := utils.GetHttpParam(w, ps, "cid")
	if err != nil {
		return
	}

	_, err = utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	// Solo l'utente che ha fatto il commento lo può rimuovere

	res, err := rt.db.UncommentPhoto(cid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrPhotoNotFound.Error()+" or "+utils.ErrCommentNotFound.Error(), http.StatusNotFound)
		return
	}

	w.Write([]byte("Comment deleted succesfully"))
}
