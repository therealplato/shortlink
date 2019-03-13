package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGolden(t *testing.T) {
	t.Run("root form", func(t *testing.T) {
		e := &endpoint{}
		s := httptest.NewServer(e)
		res, err := http.Get(s.URL)
		require.Nil(t, err)
		bb, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		golden, err := ioutil.ReadFile("testdata/root.golden.html")
		require.Nil(t, err)
		assert.Equal(t, golden, bb, string(bb))
	})
	t.Run("slug preview", func(t *testing.T) {
		e := &endpoint{}
		sv := httptest.NewServer(e)
		st := &mockStore{
			slug: "preview/abc",
			link: "https://therealplato.com",
			base: sv.URL,
			err:  nil,
		}
		e.store = st

		res, err := http.Get(sv.URL)
		require.Nil(t, err)
		bb, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		golden, err := ioutil.ReadFile("testdata/preview.golden.html")
		require.Nil(t, err)
		assert.Equal(t, golden, bb, string(bb))
	})
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

func TestLookupSlug(t *testing.T) {
	t.Run("from valid shortlink to long link", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have redirected")
	})
	t.Run("from invalid shortlink, non-url to 404", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have 404'd")
	})
	t.Run("from invalid shortlink, url to implicit creation", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have created")
	})
	t.Run("with /preview/ prefix", func(t *testing.T) {
		t.Skip("unimplemented")
		assert.True(t, false, "should have displayed both links")
	})
}
