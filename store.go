package main

import "errors"

type store interface {
	Lookup(slug string) (shortlink, error)
	Save(sl shortlink) error
}

type pqStore struct {
	URI     string
	baseURL string
}

func NewPQStore(cfg config) *pqStore {
	return &pqStore{
		URI:     cfg.PostgresURL,
		baseURL: cfg.BaseURL,
	}
}

func (s *pqStore) Validate(cfg config) error {
	// Open and ping connection to uri
	// Confirm tables exist
	// Confirm base_url is set
	return nil
}

func (s *pqStore) Lookup(slug string) (shortlink, error) {
	return shortlink{
		slug: "asdf",
		link: "hjkl",
		base: "http://localhost:8000",
	}, nil
}

func (s *pqStore) Save(sl shortlink) error {
	return errors.New("unimplemented")
}
