package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	username := r.URL.Query().Get("username")

	if !utils.HttpValidateUsername(w, username) {
		return
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
			utils.SetHeaderJson(w)
			err = json.NewEncoder(w).Encode(user)

			if err != nil {
				http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
			}
		}
	}
}
