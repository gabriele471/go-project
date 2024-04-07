package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setNewUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")

	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}

	err = r.ParseForm()
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	user.Username, err = decodeQueryParamsUsername(r)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	// missing username check to avoid same names
	err = rt.db.UpdateUsername(user.Id, user.Username)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, user.Username, http.StatusOK)
}
