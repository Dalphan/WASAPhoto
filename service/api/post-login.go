package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")

	var UID int
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get 'name' field from requestBody
	username := requestBody["name"]
	if !utils.HttpValidateUsername(w, username) {
		return
	}

	// Check if the user exists, if yes then return the user
	user, res, err := rt.db.FindUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user is not found, create a new user and return the UserID
	if res == database.NO_ROWS {
		UID, _, err = rt.db.CreateUser(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else { // If the user exists, get the UserID
		UID = user.UserID
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(UID)

	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
