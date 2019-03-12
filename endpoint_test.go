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
	t.Run("from root form post", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have created and previewed shortlink")
	})
	t.Run("from link suffix to a new destination", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have created and previewed shortlink")
	})
	t.Run("from link suffix to an existing destination", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have previewed existing shortlink")
	})
}

func TestLookup(t *testing.T) {
	t.Run("from valid shortlink to long link", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have redirected")
	})
	t.Run("from invalid shortlink to 404", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have 404'd")
	})
	t.Run("with /preview/ prefix", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have displayed both links")
	})
}
