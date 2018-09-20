package main

import (
	"net/http"
	"strings"
)

type shortlink struct {
	slug string
	link string
	base string
}

func (s shortlink) String() string {
	return strings.TrimSuffix(s.base, "/") + "/" + s.slug
}

func (s shortlink) Target() string {
	return s.link
}

func (s shortlink) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", s.String())
	w.WriteHeader(http.StatusMovedPermanently)
}
