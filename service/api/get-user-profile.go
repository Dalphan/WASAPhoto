package api

import (
	"errors"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/api/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	username := r.URL.Query().Get("username")

	err := utils.ValidateUsername(username)

	if errors.Is(err, utils.ErrUsernameMissing) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if errors.Is(err, utils.ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		//other way
		//w.WriteHeader(http.StatusNotAcceptable)
		//fmt.Fprint(w, err.Error())
	} else {
		//cerca nel db
	}
}
