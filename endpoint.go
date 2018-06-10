package main

import (
	"context"
	"log"
	"net/http"
)

type endpoint struct {
	ctx   context.Context
	store interface{}
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(root)
	if err != nil {
		log.Println(err)
	}
}
