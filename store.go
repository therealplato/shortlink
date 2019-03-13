package main

import "errors"

type store interface {
	LookupSlug(slug string) (shortlink, error)
	LookupLink(link string) (shortlink, error)
	Save(sl shortlink) error
}

type pqStore struct {
	URI     string
	baseURL string
}

var ErrNotFound = errors.New("item not found")

func NewPQStore(cfg config) *pqStore {
	return &pqStore{
		URI:     cfg.PostgresURI,
		baseURL: cfg.BaseURL,
	}
}

func (s *pqStore) Validate(cfg config) error {
	// Open and ping connection to uri
	// Confirm tables exist
	// Confirm base_url is set
	return nil
}

func (s *pqStore) LookupSlug(slug string) (shortlink, error) {
	return shortlink{
		slug: "asdf",
		link: "hjkl",
		base: "http://localhost:8000",
	}, nil
}
func (s *pqStore) LookupLink(slug string) (shortlink, error) {
	return shortlink{
		slug: "asdf",
		link: "hjkl",
		base: "http://localhost:8000",
	}, nil
}

func (s *pqStore) Save(sl shortlink) error {
	return errors.New("unimplemented")
}

type mockStore struct {
	slug string
	link string
	base string
	err  error
}

func (s *mockStore) LookupSlug(slug string) (shortlink, error) {
	return shortlink{
		slug: s.slug,
		link: s.link,
		base: s.base,
	}, s.err
}

func (s *mockStore) LookupLink(slug string) (shortlink, error) {
	return shortlink{
		slug: s.slug,
		link: s.link,
		base: s.base,
	}, s.err
}

func (s *mockStore) Save(sl shortlink) error {
	return s.err
}
