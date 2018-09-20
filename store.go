package main

import "errors"

type store interface {
	Lookup(slug string) (shortlink, error)
	Save(sl shortlink) error
}

type pqStore struct {
	uri string
}

func NewPQStore(uri string) *pqStore {
	return &pqStore{
		uri: uri,
	}
}

func (s pqStore) Lookup(slug string) (shortlink, error) {
	return shortlink{
		slug: "asdf",
		link: "hjkl",
		base: "http://localhost:8000",
	}, nil
}

func (s pqStore) Save(sl shortlink) error {
	return errors.New("unimplemented")
}
