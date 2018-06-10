package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootRendersUI(t *testing.T) {
	e := endpoint{}
	s := httptest.NewServer(e)
	res, err := http.Get(s.URL)
	assert.Nil(t, err)
}

func TestRedirects(t *testing.T) {
	t.Run("to a valid shortlink", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have redirected")
	})
	t.Run("to an invalid shortlink", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have 404'd")
	})
	t.Run("when shortlink service was the referrer", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have displayed html link rather than redirecting")
	})
	t.Run("when the path includes prefix segments", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have prefixed plato-a-")
	})
}
