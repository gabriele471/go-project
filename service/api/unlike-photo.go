package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unLikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	// this should be passed by the fe (being the name of the pic)
	// so no error handling
	postId := decodeQueryParamsPostId(r)

	err = rt.db.RemoveLike(postId, user.Id)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}

	encodeResponse(w, Msg200, http.StatusOK)

}
