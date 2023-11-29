package main

import (
	"flag"
	"github.com/daolis/training/todo-api/cmd/api"
	"log/slog"
	"os"
)

func main() {
	config := &api.AppConfig{}

	flag.StringVar(&config.Addr, "addr", ":8080", "API server port (e.g. ':8080')")
	flag.StringVar(&config.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := api.NewApp(config, log)
	err := app.StartServer()
	if err != nil {
		log.Error(err.Error())
	}
}
