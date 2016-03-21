package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"

	"github.com/crowdriff/traffic"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func bye(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, %s!", c.URLParams["name"])
}

func main() {
	// Fire up a new Goji web server with 2 handlers
	serverURL := "127.0.0.1:8000"
	server := web.New()
	server.Get("/hello/:name", hello)
	server.Get("/bye/:name", bye)
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	go func() {
		err := graceful.ListenAndServe(serverURL, server)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new traffic generator that will execute 1000 requests
	gen := traffic.NewGenerator(1000, 1)
	// Add a traffic pattern that'll hit the /hello endpoint with
	// 25% probability
	gen.AddPattern(&traffic.Pattern{
		Probability: 25,
		Fn: func() {
			URL := fmt.Sprintf("http://%s/hello/world", serverURL)
			http.Get(URL)
		},
	})
	// Add a second traffic pattern that'll hit the /bye endpoint with
	// 75% probability
	gen.AddPattern(&traffic.Pattern{
		Probability: 75,
		Fn: func() {
			URL := fmt.Sprintf("http://%s/bye/world", serverURL)
			http.Get(URL)
		},
	})
	// Execute the traffic generator
	gen.Execute()
	graceful.ShutdownNow()
}
