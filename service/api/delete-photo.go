package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dalphan/WASAPhoto/service/database"
	"github.com/Dalphan/WASAPhoto/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.SetHeaderText(w)

	pid, err := utils.GetHttpParam(w, ps, "pid")
	if err != nil {
		return
	}

	photo, res, err := rt.db.GetPhotoById(pid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if res == database.NO_ROWS {
		http.Error(w, utils.ErrPhotoNotFound.Error(), http.StatusNotFound)
		return
	}

	_, err = utils.GetAuthorization(w, r, int(photo.UserID))
	if err != nil {
		return
	}

	res, err = rt.db.DeletePhoto(pid)
	if res == database.ERROR { // The photo exists in the database, so sql.NoRows isn't checked
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//ritorna encoded la foto eliminata?
	utils.SetHeaderJson(w)
	err = json.NewEncoder(w).Encode("FOTO ELIMINATA")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
