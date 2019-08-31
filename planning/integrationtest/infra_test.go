package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/szabba/assert"
	"go.uber.org/dig"
	"gocloud.dev/docstore"
)

type NiceStatus int

func (ns NiceStatus) String() string {
	asInt := int(ns)
	return fmt.Sprintf("%s (%d)", http.StatusText(asInt), asInt)
}

func startServer(t *testing.T) *httptest.Server {
	var router *mux.Router
	c := dig.New()
	// TODO: add providers
	err := c.Invoke(func(r *mux.Router) { router = r })
	assert.That(err == nil, t.Fatalf, "cannot wire server: %s", err)
	return httptest.NewServer(router)
}

func loadFixture(t *testing.T, fixtureName string) {
	f, err := os.Open(fixtureName)
	assert.That(err == nil, t.Fatalf, "cannot load fixture %q data: %s", fixtureName, err)
	defer f.Close()

	colls := make(map[string][]interface{})
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	err = dec.Decode(&colls)
	assert.That(err == nil, t.Fatalf, "cannot parse fixture %q data: %s", fixtureName, err)

	for url, docs := range colls {
		coll, err := docstore.OpenCollection(context.Background(), url)
		assert.That(err == nil, t.Fatalf, "cannot open collection %q: %s", url, err)

		for _, d := range docs {
			err = coll.Put(context.Background(), d)
			assert.That(err == nil, t.Fatalf, "cannot fill collection %s: cannot store document %#v: %s", url, d, err)
		}
	}
}
