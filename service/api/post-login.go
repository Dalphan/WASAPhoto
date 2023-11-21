package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	username := ps.ByName("name")
	var UID int

	// Controlla se esiste l'utente e nel caso lo ritorna
	user, res, err := rt.db.FindUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Se non ha trovato niente, allora crea l'utente e ritorna l'id
	if res == database.NO_ROWS {
		UID, _, err = rt.db.CreateUser(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else { // Utente trovato, prendi lo UserID
		UID = int(user.UserID)
	}

	json.NewEncoder(w).Encode(UID)
}
