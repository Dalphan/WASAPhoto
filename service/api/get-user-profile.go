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

	// Check for Authorization
	uid, err := utils.GetAuthorization(w, r)
	if err != nil {
		return
	}

	if !utils.HttpValidateUsername(w, username) {
		return
	} else {
		// Valid username, search the profile in the database
		user, res, err := rt.db.FindUserByUsername(username, uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// User not found, return 404 Not Found
		if res == database.NO_ROWS {
			http.Error(w, utils.ErrUserNotFound.Error(), http.StatusNotFound)
		} else { // User found

			// Check if the User has banned the one searching
			if CheckBanned(w, rt, user.UserID, uid, utils.ErrUserNotFound) {
				return
			}

			utils.SetHeaderJson(w)
			err = json.NewEncoder(w).Encode(user)

			if err != nil {
				http.Error(w, utils.ErrEncodingJson.Error(), http.StatusInternalServerError)
			}
		}
	}
}
