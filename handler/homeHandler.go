package handler

import "net/http"

type homeHandler struct {
}

//NewHomeHandler app
func NewHomeHandler() IHandler {
	return homeHandler{}
}

func (h homeHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("home"))
}
