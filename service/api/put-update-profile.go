package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	var user utils.User
	err := json.NewDecoder(r.Body).Decode(&user)

	//Error getting the user from the requestBody
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Check username validity
	err = utils.ValidateUsername(user.Username)
	if errors.Is(err, utils.ErrUsernameMissing) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errors.Is(err, utils.ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

}
