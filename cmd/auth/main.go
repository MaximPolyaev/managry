package main

import (
	"os"

	"github.com/MaximPolyaev/managry/internal/adapters/authserver"
	"github.com/MaximPolyaev/managry/internal/handlers/authhandler"
	"github.com/MaximPolyaev/managry/internal/utils/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		authHandler := authhandler.New()

		r.Post("/signin", authHandler.SignInHandleFunc)
		r.Post("/signup", authHandler.SignUpHandleFunc)
	})

	serv := authserver.New(log, r)

	return serv.ListenAndServe()
}
