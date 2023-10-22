package main

import (
	"os"

	"managry-go/internal/adapters/authserver"
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
	serv := authserver.New(log)

	return serv.ListenAndServe()
}
