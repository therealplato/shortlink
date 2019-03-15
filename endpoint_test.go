package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Golden output is populated by make bake
func TestGolden(t *testing.T) {
	e := &endpoint{}
	sv := httptest.NewServer(e)
	t.Run("root form", func(t *testing.T) {
		res, err := http.Get(sv.URL)
		require.Nil(t, err)
		bb, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		golden, err := ioutil.ReadFile("testdata/root.golden.html")
		require.Nil(t, err)
		assert.Equal(t, golden, bb, string(bb))
	})
	t.Run("slug preview", func(t *testing.T) {
		st := &mockStore{
			slug: "abc",
			link: "http://therealplato.com",
			err:  nil,
		}
		e.store = st
		e.baseURL = "http://localhost:8000/"

		res, err := http.Get(sv.URL + "/preview/abc")
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
		st := &mockStore{
			slug: "abc",
			link: "https://duck.com",
		}
		e := &endpoint{store: st, baseURL: "https://base.url/still/"}
		req, err := http.NewRequest(http.MethodGet, "https://base.url/still/abc", nil)
		require.Nil(t, err)
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusFound, rr.Code, "should be %v", http.StatusFound)
		assert.Equal(t, st.link, rr.Header().Get("Location"), "should have Location header")
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
