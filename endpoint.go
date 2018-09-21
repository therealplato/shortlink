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

var rxSlug = regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
var rxNotSlug = regexp.MustCompile(`[^a-zA-Z0-9\-]`)
var rxDashes = regexp.MustCompile(`-+`) // repeated dashes

func slugify(s string) string {
	s = rxNotSlug.ReplaceAllString(s, "-")
	s = rxDashes.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
