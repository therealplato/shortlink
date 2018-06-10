package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/therealplato/lifecycle"
)

type postgresStore struct {
	DSN string
}

func main() {
	l := lifecycle.Begin()
	go l.HealthCheck(os.Getenv("HEALTHCHECK_LISTEN_ADDR"))
	e := &endpoint{
		ctx: l.Ctx,
		store: postgresStore{
			DSN: os.Getenv("POSTGRES_DSN"),
		},
	}

	s := http.Server{
		Addr:    os.Getenv("SERVER_LISTEN_ADDR"),
		Handler: e,
	}
	go func() {
		fmt.Printf("i live to serve shortlinks on %s\n", s.Addr)
		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-l.Ctx.Done()
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	s.Shutdown(ctx2)
}
