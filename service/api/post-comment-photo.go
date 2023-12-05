package api

import (
	"encoding/json"
	"net/http"

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

	uid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	//CONTROLLA CHE IL COMMENTO E' DELL'UTENTE

	comment.PhotoID = uint(pid)
	comment.UserID = uint(uid)
	comment.Date = utils.NowFormat()

}
