package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreLooksUpShortlink(t *testing.T) {
	s := store{}
	expected := "http://localhost:8000/asdf"
	link := s.Lookup("asdf")
	assert.Equal(t, expected, link.String())
}

func TestStoreSavesShortlink(t *testing.T) {
	s := store{}
	sl := shortlink{slug: "asdf"}
	err := s.Save(sl)
	assert.Nil(t, err)
	assert.Equal(t, "asdf", sl.slug)
	assert.Equal(t, "https://base/asdf", sl.String())
	assert.Equal(t, "https://src/target", sl.Target())
}
