package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.uber.org/dig"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
	"gocloud.dev/server"

	"github.com/szabba/upstep/planning"
	docstoreadapters "github.com/szabba/upstep/planning/docstore"
	"github.com/szabba/upstep/planning/handler"
)

func main() {
	err := run()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func run() error {
	c := dig.New()

	c.Provide(server.New)
	c.Provide(defaultServerOptions)
	c.Provide(handler.NewMux)
	c.Provide(handler.NewInit)
	c.Provide(handler.NewPlan)
	c.Provide(newPlannerRepository)
	c.Provide(newPlanRepository)
	c.Provide(newStepRepository)

	return c.Invoke(runServer)
}

func defaultServerOptions() *server.Options { return nil }

func runServer(srv *server.Server) error {
	port := envOrElse("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s", addr)
	return srv.ListenAndServe(addr)
}

func newPlannerRepository() (planning.PlannerRepository, error) {
	ctx := context.Background()
	url := envOrElse("PLANNER_COLLECTION_URL", "mem://planner/ID")
	coll, err := docstore.OpenCollection(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q: %s", url, err)
	}
	return docstoreadapters.NewPlannerRepository(coll), nil
}

func newPlanRepository() (planning.PlanRepository, error) {
	ctx := context.Background()
	url := envOrElse("PLAN_COLLECTION_URL", "mem://plan/ID")
	coll, err := docstore.OpenCollection(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q: %s", url, err)
	}
	return docstoreadapters.NewPlanRepository(coll), nil
}

func newStepRepository() (planning.StepRepository, error) {
	ctx := context.Background()
	url := envOrElse("STEP_COLLECTION_URL", "mem://step/ID")
	coll, err := docstore.OpenCollection(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q: %s", url, err)
	}
	return docstoreadapters.NewStepRepository(coll), nil
}

func envOrElse(name, whenEmpty string) string {
	v := os.Getenv(name)
	if v == "" {
		return whenEmpty
	}
	return v
}
