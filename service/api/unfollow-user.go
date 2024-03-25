package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	username, err := decodeQueryParamsUsername(r)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	targetUser, err := rt.isUserRegistered("", username)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	err = rt.db.RemoveFollow(user.Id, targetUser.Id)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, Msg200, http.StatusOK)
}
