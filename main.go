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

type endpoint struct {
	ctx   context.Context
	store interface{}
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("adsf"))
}

type postgresStore struct {
	DSN string
}

func main() {
	fmt.Println("hello")
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
		fmt.Printf("serving shortlinks on %s\n", s.Addr)
		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	fmt.Println("i live to serve")
	<-l.Ctx.Done()
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	s.Shutdown(ctx2)
}
