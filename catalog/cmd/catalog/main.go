package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/tinrab/spidey/catalog"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	cfg.DatabaseURL = "postgres://postgres:postgres@localhost/spidey?sslmode=disable"
	var r catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		fmt.Println("DB" + cfg.DatabaseURL)
		r, err = catalog.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := catalog.NewService(r)
	log.Fatal(catalog.ListenGRPC(s, 8080))
}
