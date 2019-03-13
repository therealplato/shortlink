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
	// Explicit creation:
	if r.Method == http.MethodPost {
		e.create(w, r)
		return
	}
	// Landing:
	if strings.TrimLeft(r.URL.Path, "/") == "" {
		_, err := w.Write(root)
		if err != nil {
			log.Println(err)
		}
		return
	}
	// Explicit preview:
	if strings.HasPrefix(r.URL.Path, "/preview") {
		e.preview(w, r)
		return
	}
	// Slug lookup or implict creation:
	e.slug(w, r)
}

// slug may be invoked with a known slug, an unknown slug, or a url to shorten
func (e *endpoint) slug(w http.ResponseWriter, r *http.Request) {
	var (
		input = strings.TrimLeft(r.URL.Path, "/")
		isURL = rxProbableURL.MatchString(input)
		sl    shortlink
		err   error
	)
	if isURL {
		e.create(w, r)
		return
	}
	sl, err = e.store.LookupSlug(input)
	if err == ErrNotFound {
		handleNotFound(w)
		return
	}
	if err != nil {
		log.Printf("slug error: %q\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, sl.link, http.StatusFound)
}

func (e *endpoint) create(w http.ResponseWriter, r *http.Request) {
	slug := slugify(r.URL.Path)

	// it might come in path instead of form
	link := r.FormValue("link")
	if link == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("requires valid URL as 'link' form parameter"))
		return
	}
	//
	// sl = shortlink{
	// 	slug: randomSlug(),
	// 	link: input,
	// }
	// err = e.store.Save(sl)

	fmt.Printf("Linking %q to %q\n", link, e.cfg.BaseURL+slug)
	// parse form to get destination url
	// generate slug
	// instantiate shortlink
	// persist
	return
}

func (e *endpoint) preview(w http.ResponseWriter, r *http.Request) {
	// lookup slug maybe 404
}

func handleNotFound(w http.ResponseWriter) {
	m := excuses[rand.Intn(len(excuses))]
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(m))
}

var rxSlug = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
var rxNotSlug = regexp.MustCompile(`[^a-zA-Z0-9\-]`)
var rxProbableURL = regexp.MustCompile(`^https?://`)
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
