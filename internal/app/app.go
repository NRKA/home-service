package app

import (
	"context"
	"github.com/NRKA/home-service/internal/server/router"
	"github.com/NRKA/home-service/internal/service/auth"
	"github.com/NRKA/home-service/internal/service/flat"
	"github.com/NRKA/home-service/internal/service/house"
	"github.com/NRKA/home-service/internal/service/sender"
	"github.com/NRKA/home-service/pkg/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

const port = "PORT"

func Run(config postgres.DatabaseConfig) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	db, err := postgres.NewDB(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer db.GetPool().Close()

	auth := auth.NewHandler(db)
	house := house.NewHandler(db)
	flat := flat.NewHandler(db)
	s := sender.NewHandler(db)
	r := router.New(auth, house, flat, s)

	port := os.Getenv(port)
	var wg sync.WaitGroup
	wg.Add(2)

	sender := sender.New()
	go func() {
		defer wg.Done()

		if err := sender.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err = http.ListenAndServe(port, r); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	wg.Wait()
}
