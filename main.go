package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aminkhalilii/go-ticket/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//loads
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("can not load config")
	}
	if config.APPDEBUG == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	dns := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DBUSERNAME, config.DBPASSWORD, config.DBHOST, config.DBPORT, config.DBDATABASE)
	dbpool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

}
