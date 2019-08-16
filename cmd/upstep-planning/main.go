package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/dig"
	"gocloud.dev/server"
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
	c.Provide(func() *server.Options { return nil })
	c.Provide(func() http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("responding to %s", r.RequestURI)
			fmt.Fprintf(w, "Hello from %s", r.RequestURI)
		})
	})
	return c.Invoke(func(srv *server.Server) error {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		addr := fmt.Sprintf(":%s", port)
		log.Printf("listening on %s", addr)
		return srv.ListenAndServe(addr)
	})
}
