package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortlink(t *testing.T) {
	sl := shortlink{
		slug: "a",
		link: "https://b",
		base: "https://base/",
	}
	assert.Equal(t, "https://base/a", sl.String())
	assert.Equal(t, "https://b", sl.Target())
}
