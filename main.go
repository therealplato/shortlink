package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	rand.Seed(time.Now().UnixNano())

	st := NewPQStore(cfg)
	st.MustConn()
	e := &endpoint{
		ctx:     l.Ctx,
		store:   st,
		baseURL: cfg.BaseURL,
	}

	sv := http.Server{
		Addr:    cfg.ShortlinkAddr,
		Handler: e,
	}
	go func() {
		fmt.Printf("i live to serve shortlinks on %s\n", sv.Addr)
		err := sv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-l.Ctx.Done()
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	sv.Shutdown(ctx2)
}
