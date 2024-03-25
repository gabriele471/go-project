package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	commentId := decodeQueryParamsCommentId(r)

	err = rt.db.RemoveComment(postId, user.Id, commentId)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, Msg200, http.StatusOK)
}
