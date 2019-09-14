package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/szabba/assert"
	"go.uber.org/dig"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/memdocstore"

	"github.com/szabba/upstep/planning"
	docstoreglue "github.com/szabba/upstep/planning/docstore"
	"github.com/szabba/upstep/planning/httpglue"
)

type NiceStatus int

func (ns NiceStatus) String() string {
	asInt := int(ns)
	return fmt.Sprintf("%s (%d)", http.StatusText(asInt), asInt)
}

func startServer(t *testing.T) *httptest.Server {
	var router *mux.Router
	c := dig.New()
	c.Provide(newPlanRepository)
	c.Provide(httpglue.NewPlans)
	c.Provide(httpglue.NewMux)
	err := c.Invoke(func(r *mux.Router) { router = r })
	assert.That(err == nil, t.Fatalf, "cannot wire server: %s", err)
	return httptest.NewServer(router)
}

func newPlanRepository() (repo planning.PlanRepository, _ error) {
	ctx := context.Background()
	coll, err := docstore.OpenCollection(ctx, "mem://plans/ID")
	if err != nil {
		return nil, fmt.Errorf("cannot create %v: %w", reflect.TypeOf(&repo).Elem(), err)
	}
	return docstoreglue.NewPlanRepository(coll), nil
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
