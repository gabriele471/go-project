package api

import (
	"io/ioutil"
	"myproject/service/types"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	authorizationHeader := r.Header.Get("Authorization")
	user, err := rt.isTokenValid(authorizationHeader)
	if err != nil {
		encodeResponse(w, Msg401, http.StatusUnauthorized)
		return
	}
	// image is sent within the header
	// later saved in a dir, named as its Post Id
	// Determine the file extension based on the Content-Type header
	PhotoId := generateGenericToken()
	// downloadDirectory := "/service/images" // relative path from service/api/filethatsavesit
	downloadDirectory := "./service/images"

	path := filepath.Join(downloadDirectory, PhotoId+".jpg")
	_, err = os.Create(path)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}

	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		encodeResponse(w, Msg400, http.StatusBadRequest)
		return
	}
	err = saveResponseBodyToFile(responseBody, path)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}
	post := types.NewPost(PhotoId, user.Id, path, user.Username)
	post.OwnerUsername = user.Username
	err = rt.db.InsertPost(*post)
	if err != nil {
		encodeResponse(w, Msg500, http.StatusInternalServerError)
		return
	}

	encodeResponse(w, Msg200+", Created", http.StatusCreated)
}
