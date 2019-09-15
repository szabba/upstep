package integrationtest

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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
