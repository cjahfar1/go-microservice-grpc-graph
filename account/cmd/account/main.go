package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"github.com/tinrab/spidey/account"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	cfg.DatabaseURL = "postgres://postgres:postgres@localhost/spidey?sslmode=disable"
	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		fmt.Println("DB URL:" + cfg.DatabaseURL)
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
