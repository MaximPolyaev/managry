package signinhandler

import "net/http"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("sign in"))
}
