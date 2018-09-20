package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/therealplato/lifecycle"
)

func main() {
	log.Println("loading config...")
	cfg := MustLoadConfig()
	log.Println("config loaded")
	l := lifecycle.Begin()
	go l.HealthCheck(cfg.HealthcheckAddr)
	e := &endpoint{
		ctx:   l.Ctx,
		store: NewPQStore(cfg),
	}

	s := http.Server{
		Addr:    cfg.ShortlinkAddr,
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
