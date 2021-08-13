package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "inventory-api", log.LstdFlags)

	// create mux
	router := mux.NewRouter()

	// create handler
	ih := newInventoryHandler(l)

	router.HandleFunc("/", ih.listInventory)

	s := http.Server{
		Addr:     ":3001",
		Handler:  router,
		ErrorLog: l,
	}

	// start server
	go func() {
		l.Println("Starting server on port 3001")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until signal recieved
	sig := <-c
	log.Println("Got signal: ", sig)

	// graceful shutdown, max 2min until forced close
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	s.Shutdown(ctx)
}
