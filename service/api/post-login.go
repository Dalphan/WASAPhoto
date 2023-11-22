package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	bytedata, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var UID int
	var result map[string]string
	json.Unmarshal([]byte(bytedata), &result)
	username := result["name"]

	fmt.Println("USername: ", username)

	// Controlla se esiste l'utente e nel caso lo ritorna
	user, res, err := rt.db.FindUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(res)

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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UID)
}
