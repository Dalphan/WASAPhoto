package api

import (
	"errors"
	"net/http"

	_ "github.com/Dalphan/WASAPhoto/service/api/utilities"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	username := r.URL.Query().Get("username")

	err := validateUsername(username)

	if errors.Is(err, ErrUsernameMissing) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if errors.Is(err, ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		//other way
		//w.WriteHeader(http.StatusNotAcceptable)
		//fmt.Fprint(w, err.Error())
	} else {
		//cerca nel db
	}
}
