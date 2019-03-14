package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPQStoreLookupSlug(t *testing.T) {
	t.Skip()
	s := pqStore{}
	expected := "http://localhost:8000/asdf"
	link, err := s.LookupSlug("asdf")
	assert.Nil(t, err)
	assert.Equal(t, expected, link.String())
}

func TestStoreSavesShortlink(t *testing.T) {
	t.Skip()
	t.Run("with missing link", func(t *testing.T) {
		s := pqStore{}
		sl := shortlink{slug: "asdf"}
		err := s.Save(sl)
		assert.NotNil(t, err)
	})
	t.Run("with present link", func(t *testing.T) {
		s := pqStore{}
		sl := shortlink{slug: "asdf", link: "https://duckduckgo.com"}
		err := s.Save(sl)
		assert.NotNil(t, err)
	})
}
