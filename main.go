package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/fenpaws/energeticon/handlers"
	_ "github.com/go-mojito/defaults"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/log"
	"github.com/go-mojito/mojito/middleware"
	"strconv"
)

type ConfigDatabase struct {
	Port int `env:"PORT" envDefault:"8123"`
}

func main() {

	cfg := ConfigDatabase{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Register application routers...")

	mojito.GET("/", handlers.HomeHandler)
	mojito.WithMiddleware(middleware.RequestLogger)

	log.Infof("Server has started on %s", "0.0.0.0:"+strconv.Itoa(cfg.Port))
	log.Error(mojito.ListenAndServe("0.0.0.0:" + strconv.Itoa(cfg.Port)))
}
