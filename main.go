package main

import (
	"os"

	"github.com/aminkhalilii/go-ticket/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//loads
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("can nnot load config")
	}
	if config.APPDEBUG == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

}
