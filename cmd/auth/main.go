package main

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"managry-go/internal/adapters/authserver"
	"managry-go/internal/handlers/signinhandler"
	"managry-go/internal/handlers/signuphandler"
	"managry-go/internal/utils/logger"
)

func main() {
	log := logger.New(os.Stdout)

	err := run(log)
	if err != nil {
		log.Fatal(err)
	}
}

func run(log *logger.Logger) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		signIn := signinhandler.New()
		signUp := signuphandler.New()

		r.Post("/signin", signIn.Handler)
		r.Post("/signup", signUp.Handler)
	})

	serv := authserver.New(log, r)

	return serv.ListenAndServe()
}
