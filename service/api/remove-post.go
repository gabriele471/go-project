package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	postId := decodeQueryParamsPostId(r)
	err = rt.db.RemovePost(user.Id, postId)
	if err != nil {
		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}

	imagesDir := "./webui/public/images"
	fileToDelete := postId + ".jpg"

	filePathToDelete := filepath.Join(imagesDir, fileToDelete)

	err = os.Remove(filePathToDelete)
	if err != nil {
		rt.baseLogger.Error(err)
	}
	encodeResponse(w, Msg200, http.StatusOK)

}
