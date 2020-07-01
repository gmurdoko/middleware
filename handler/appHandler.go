package handler

import "net/http"

//IHandler app
type IHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}
