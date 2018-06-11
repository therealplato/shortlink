package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortlinkIsURLSafe(t *testing.T) {
	sl := New("hjkl")
	u := url.URL{Path: sl.slug}
	enc := u.EscapedPath()
	assert.Equal(t, sl.slug, enc)
}
