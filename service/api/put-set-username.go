package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	uid, err := utils.GetHttpParam(w, ps, "id")
	if err != nil { //Error getting the user ID
		return
	}

	var requestBody map[string]string
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get 'username' field from requestBody
	username := requestBody["name"]
	if !utils.HttpValidateUsername(w, username) {
		return
	}

	user, res, err := rt.db.UpdateUsername(uid, username)
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
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
	}
}
