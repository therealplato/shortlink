package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
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
		e.createPost(w, r)
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
		e.createGet(w, r)
		return
	}
	sl, err = e.store.LookupSlug(input)
	if err == ErrNotFound {
		handleNotFound(w)
		return
	}
	if err != nil {
		log.Printf("slug: %q error: %q\n", input, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, sl.link, http.StatusFound)
}

// receie a POST with slug and link
func (e *endpoint) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	sl := shortlink{}
	err := r.ParseForm()
	if err != nil {
		handleNotFound(w)
		return
	}
	sl.slug = r.PostFormValue("slug")
	sl.link = r.PostFormValue("link")
	if sl.slug == "" {
		sl.slug = randomSlug()
	}
	if sl.link == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing 'link' form parameter"))
		return
	}
	_, err = url.Parse(sl.link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("link did not parse as url"))
		return
	}

	for {
		err := e.store.Save(sl)
		switch err {
		case nil:
			break
		case ErrDupe:
			sl.slug = randomSlug()
			continue
		default:
			log.Printf("saving failed: %#v", sl)
			handleNotFound(w)
			return
		}
	}
	fmt.Printf("%q -> %q\n", sl.slug, sl.link)
	http.Redirect(w, r, "/preview/"+sl.slug, http.StatusFound)
	return
}
func (e *endpoint) createGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	sl := shortlink{
		slug: randomSlug(),
		link: strings.TrimLeft(r.URL.Path, "/"),
	}
	_, err := url.Parse(sl.link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("link did not parse as url"))
		return
	}

	for {
		err := e.store.Save(sl)
		switch err {
		case nil:
			break
		case ErrDupe:
			sl.slug = randomSlug()
			continue
		default:
			log.Printf("saving failed: %#v", sl)
			handleNotFound(w)
			return
		}
	}
	fmt.Printf("%q -> %q\n", sl.slug, sl.link)
	http.Redirect(w, r, "/preview/"+sl.slug, http.StatusFound)
	return
}

func (e *endpoint) preview(w http.ResponseWriter, r *http.Request) {
	log.Printf("preview %q\n", r.URL.Path)
	slug := strings.TrimLeft(r.URL.Path, "/preview/")
	sl, err := e.store.LookupSlug(slug)
	if err != nil {
		handleNotFound(w)
		return
	}
	fmt.Fprintf(w, "%q -> %q\n", sl.slug, sl.link)
}

func handleNotFound(w http.ResponseWriter) {
	m := excuses[rand.Intn(len(excuses))]
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(m))
}

func randomSlug() string {
	a := fragments[rand.Intn(len(fragments))]
	b := fragments[rand.Intn(len(fragments))]
	c := fragments[rand.Intn(len(fragments))]
	return fmt.Sprintf("%s%s%s", a, b, c)
}

/*
func slugify(s string) string {
	s = rxNotSlug.ReplaceAllString(s, "-")
	s = rxDashes.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
// TODO: have store give fresh slugs
*/
