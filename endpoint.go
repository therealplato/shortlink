package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
)

type endpoint struct {
	ctx   context.Context
	store store
	cfg   config
}

// NewEndpoint instantiates a shortlink endpoint
func NewEndpoint(ctx context.Context, cfg config, store store) *endpoint {
	e := &endpoint{
		ctx:   ctx,
		store: store,
		cfg:   cfg,
	}
	return e
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		e.create(w, r)
		return
	}
	switch r.URL.Path {
	case "/", "":
		_, err := w.Write(root)
		if err != nil {
			log.Println(err)
		}
		return
	default:
		e.slug(w, r)
		return
	}
}

func (e *endpoint) slug(w http.ResponseWriter, r *http.Request) {
	// lookup slug maybe 404
	// if referrer is self, display a friendly shortlink info page
	// else redirect
	// sl, err := e.store.Lookup(r.URL.Path)
	// if err == sql.ErrNoRows {
	// 	handleNotFound(w)
	// }
	// sl.Handle(w, r)
}

func (e *endpoint) create(w http.ResponseWriter, r *http.Request) {
	slug := slugify(r.URL.Path)

	link := r.FormValue("link")
	if link == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("requires valid URL as 'link' form parameter"))
		return
	}

	fmt.Printf("Linking %q to %q\n", link, e.cfg.BaseURL+slug)
	// parse form to get destination url
	// generate slug
	// instantiate shortlink
	// persist
	return
}

func handleNotFound(w http.ResponseWriter) {
	m := excuses[rand.Intn(len(excuses))]
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(m))
}

var rxSlug = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
var rxNotSlug = regexp.MustCompile(`[^a-zA-Z0-9\-]`)
var rxDashes = regexp.MustCompile(`-+`) // repeated dashes

func slugify(s string) string {
	s = rxNotSlug.ReplaceAllString(s, "-")
	s = rxDashes.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func randomSlug() string {
	a := fragments[rand.Intn(len(fragments))]
	b := fragments[rand.Intn(len(fragments))]
	c := fragments[rand.Intn(len(fragments))]
	return fmt.Sprintf("%s-%s-%s", a, b, c)
}
