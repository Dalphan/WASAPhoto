package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
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
		//Lo username è valido, cerca nel database
		user, res, err := rt.db.FindUserByUsername(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Se non ha trovato niente, allora ritorna codice 404 Not Found
		if res == database.NO_ROWS {
			http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		} else { // Utente trovato, ritornalo
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	}
}
