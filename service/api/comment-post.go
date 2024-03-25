package api

import (
	"io/ioutil"
	"myproject/service/types"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	// commment content is passed as raw text in the body request
	var comment types.Comment
	comment.Username = user.Username
	comment.Message = string(body)

	err = rt.db.InsertComment(generateGenericToken(), postId, user.Id, comment.Message)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	encodeResponse(w, Msg200, http.StatusOK)

}
