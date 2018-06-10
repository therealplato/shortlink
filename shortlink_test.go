package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortlinkIsURLSafe(t *testing.T) {
	s := &shortlink{}
	l := s.New()
	u := url.URL{Path: l}
	enc := u.EscapedPath()
	assert.Equal(t, l, enc)
}
