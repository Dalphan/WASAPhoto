package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
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
	if !utils.HttpValidateUsername(w, user.Username) {
		return
	}

	//Update user in database. If the username or the user doesn't exist, is already taken it should return an error
	user, res, err := rt.db.UpdateUser(user)
	if res == database.UNIQUE_FAILED { // The updated username is already taken
		http.Error(w, utils.ErrUsernameTaken.Error(), http.StatusNotAcceptable)
		return
	}
	if res == database.NO_ROWS { // The UserID doesn't exist
		http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		return
	}
	if err != nil { // Generic error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Return the user updated in the response
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}