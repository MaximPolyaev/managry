package authserver

import (
	"net/http"
)

const defaultAddress = "localhost:8080"

type AuthServer struct {
	http.Server
	log logger
}

type logger interface {
	Info(args ...any)
}

func New(log logger, h http.Handler) *AuthServer {
	return &AuthServer{
		Server: http.Server{
			Addr:    defaultAddress,
			Handler: h,
		},
		log: log,
	}
}

func (srv *AuthServer) ListenAndServe() error {
	srv.log.Info("start server: " + defaultAddress)

	return srv.Server.ListenAndServe()
}
