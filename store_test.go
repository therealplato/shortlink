package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreLooksUpShortlink(t *testing.T) {
	s := pqStore{}
	expected := "http://localhost:8000/asdf"
	link, err := s.Lookup("asdf")
	assert.Nil(t, err)
	assert.Equal(t, expected, link.String())
}

func TestStoreSavesShortlink(t *testing.T) {
	s := pqStore{}
	sl := shortlink{slug: "asdf"}
	err := s.Save(sl)
	assert.Nil(t, err)
	assert.Equal(t, "asdf", sl.slug)
	assert.Equal(t, "https://base/asdf", sl.String())
	assert.Equal(t, "https://src/target", sl.Target())
}
