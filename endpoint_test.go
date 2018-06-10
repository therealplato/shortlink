package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootRendersUI(t *testing.T) {
	e := &endpoint{}
	s := httptest.NewServer(e)
	res, err := http.Get(s.URL)
	require.Nil(t, err)
	bb, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	golden, err := ioutil.ReadFile("testdata/root.golden.html")
	assert.Nil(t, err)
	assert.Equal(t, golden, bb, string(bb))
}

func TestCreation(t *testing.T) {
	t.Run("to a new destination", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have created")
	})
	t.Run("to an existing destination", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have redirected")
	})
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
