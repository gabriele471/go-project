package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	postId := decodeQueryParamsPostId(r)

	err = rt.db.InsertLike(generateGenericToken(), postId, *user)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	encodeResponse(w, Msg200, http.StatusOK)
}
