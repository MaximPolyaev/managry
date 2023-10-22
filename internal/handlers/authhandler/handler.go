package authhandler

import "net/http"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SignInHandleFunc(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("sign in"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignUpHandleFunc(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("sign up"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
