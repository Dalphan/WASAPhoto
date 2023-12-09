package api

import (
	"net/http"
)

func CheckBanned(w http.ResponseWriter, rt *_router, uid int, bid int, bannedErr error) bool {
	banned, err := rt.db.CheckIfBanned(uid, bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	if banned {
		http.Error(w, bannedErr.Error(), http.StatusNotFound)
		return true
	}
	return false
}
