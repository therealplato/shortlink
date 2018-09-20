package main

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"net/http"
)

type endpoint struct {
	ctx   context.Context
	store store
	mux   *http.ServeMux
}

// NewEndpoint instantiates a shortlink endpoint
func NewEndpoint(ctx context.Context, store store) *endpoint {
	e := &endpoint{
		ctx:   ctx,
		store: store,
		mux:   http.NewServeMux(),
	}

	e.mux.HandleFunc("/create", e.create)
	e.mux.HandleFunc("/", e.root)
	return e
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.mux.ServeHTTP(w, r)
}

func (e *endpoint) root(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "":
		_, err := w.Write(root)
		if err != nil {
			log.Println(err)
		}
		return
	default:
		sl, err := e.store.Lookup(r.URL.Path)
		if err == sql.ErrNoRows {
			handleNotFound(w)
		}
		sl.Handle(w, r)
	}
}

func (e *endpoint) create(w http.ResponseWriter, r *http.Request) {
	// parse form to get destination url
	// generate slug
	// instantiate shortlink
	// persist
	return
}

func handleNotFound(w http.ResponseWriter) {
	mm := []string{
		"excuse me waiter, there's a 404 in my link",
		"404 outlook not so good",
		"do you like pina coladas, and getting FILE NOT FOUND?",
		"always, I know, you'll be, 404",
		"it's a bittersweet symphony, this 404",
		"hey now, you're a rockstar, get the show on, 404",
		"sorry girl, but you missed out, well tough luck that 404's now",
		"music hits me, so hard, makes me say, 404",
		"i'm blue, 404044, 404, 404, 404044 (4 4 4..)",
		"if you wanna be my lover, you gotta 404 my friends",
		"if I 404 will you still call me superman?",
		"oops I did it again",
		"you better lose yourself in the music, the moment, you hold it, you better never 404",
		"i feel stupid, and contagious, 404 now, imitate us",
	}
	m := mm[rand.Intn(len(mm))]
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(m))
}
