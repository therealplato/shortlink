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
	ctx     context.Context
	store   store
	baseURL string
}

// WithoutBaseURL returns r.URL.String() without the base url and without leading slash
func withoutBaseURL(baseURL string, r *http.Request) string {
	suffix := strings.TrimPrefix(r.URL.String(), baseURL)
	suffix = strings.TrimLeft(suffix, "/")
	return suffix
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %q\n", r.Method, r.URL.String())
	// Explicit creation:
	if r.Method == http.MethodPost {
		e.createPost(w, r)
		return
	}

	// Landing:
	suffix := withoutBaseURL(e.baseURL, r)
	if suffix == "" {
		_, err := w.Write(root)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Explicit preview:
	if strings.HasPrefix(suffix, "preview") {
		e.preview(w, r)
		return
	}

	// Slug lookup or implict creation:
	e.slug(w, r)
}

// slug may be invoked with a known slug, an unknown slug, or a url to shorten
func (e *endpoint) slug(w http.ResponseWriter, r *http.Request) {
	var (
		suffix = withoutBaseURL(e.baseURL, r)
		isURL  = rxProbableURL.MatchString(suffix)
		sl     shortlink
		err    error
	)
	if isURL {
		e.createGet(w, r)
		return
	}
	sl, err = e.store.LookupSlug(suffix)
	switch err {
	case ErrNotFound:
		handleNotFound(w)
		return
	case nil:
		break
	default:
		log.Printf("slug: %q error: %q\n", suffix, err)
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

retry:
	for {
		err := e.store.Save(sl)
		switch err {
		case nil:
			break retry
		case ErrDupe:
			// TODO: depending on updatesAllowed, update or redirect to current
			sl.slug = randomSlug()
			continue
		default:
			log.Printf("saving failed: %#v", sl)
			handleNotFound(w)
			return
		}
	}
	fmt.Printf("%q -> %q\n", sl.slug, sl.link)
	http.Redirect(w, r, e.baseURL+"preview/"+sl.slug, http.StatusFound)
	return
}

func (e *endpoint) createGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var (
		suffix = withoutBaseURL(e.baseURL, r)
		sl     = shortlink{
			slug: randomSlug(),
			link: suffix,
		}
	)

	_, err := url.Parse(sl.link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("link did not parse as url"))
		return
	}

retry:
	for {
		err := e.store.Save(sl)
		switch err {
		case nil:
			break retry
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
	http.Redirect(w, r, e.baseURL+"preview/"+sl.slug, http.StatusFound)
	return
}

func (e *endpoint) preview(w http.ResponseWriter, r *http.Request) {
	suffix := withoutBaseURL(e.baseURL, r)
	suffix = strings.TrimPrefix(suffix, "preview/")
	sl, err := e.store.LookupSlug(suffix)
	if err != nil {
		handleNotFound(w)
		return
	}
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html><body>
<a href="%s%s">%s%s</a>
<br> -> <br>
<a href="%s">%s</a>
</body></html>
`, e.baseURL, sl.slug,
		e.baseURL, sl.slug,
		sl.link, sl.link)
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
