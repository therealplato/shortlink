package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type store interface {
	LookupSlug(slug string) (shortlink, error)
	LookupLink(link string) (shortlink, error)
	Save(sl shortlink) error
}

type pqStore struct {
	URI     string
	baseURL string
	db      *sql.DB
}

var ErrNotFound = errors.New("item not found")
var ErrDupe = errors.New("unique identifier collision")

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

func (s *pqStore) MustConn() {
	var (
		err error
	)
	timeout := time.After(30 * time.Second)
	if s.db != nil {
		err := s.db.Ping()
		if err == nil {
			return
		}
	}
connect:
	for {
		select {
		case <-timeout:
			log.Fatal("timeout connecting to postgres")

		default:
			s.db, err = sql.Open("postgres", s.URI)
			if err != nil {
				continue
			}
			err = s.db.Ping()
			if err != nil {
				continue
			}
			break connect
		}
	}

	return
}

func (s *pqStore) LookupSlug(slug string) (shortlink, error) {
	s.MustConn()
	row := s.db.QueryRow(`SELECT slug, link from shortlink where slug=$1`, slug)
	sl := shortlink{}
	err := row.Scan(&sl.slug, &sl.link)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name(), err.Code)
	}
	if err == sql.ErrNoRows {
		return sl, ErrNotFound
	}
	return sl, err
}

func (s *pqStore) LookupLink(link string) (shortlink, error) {
	s.MustConn()
	row := s.db.QueryRow(`SELECT slug, link from shortlink where link=$1`, link)
	sl := shortlink{}
	err := row.Scan(&sl.slug, &sl.link)
	return sl, err
}

func (s *pqStore) Save(sl shortlink) error {
	s.MustConn()
	_, err := s.db.Exec(`INSERT INTO shortlink (slug, link) VALUES ($1, $2)`, sl.slug, sl.link)
	// _, err := s.db.Exec(`INSERT INTO shortlink (slug, link) VALUES ($1, $2) ON CONFLICT UPDATE link=link RETURNING (slug, link)`, sl.slug, sl.link)
	return err
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
