package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rusdiyanto82/my-lessons/microservices/driver-api/handlers"
)

func main() {

	l := log.New(os.Stdout, "driver-api", log.LstdFlags)

	//create the handlers
	dr := handlers.NewDriver(l)

	//create a new servemux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", dr)

	//create a new server
	s := &http.Server{
		Addr:         ":9090",           //configure the bind address
		Handler:      sm,                //set the default handlers
		ErrorLog:     l,                 //set the logger for the server
		IdleTimeout:  120 * time.Second, //max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   //max time to read request from the client
		WriteTimeout: 1 * time.Second,   //max time to write response to the client
	}

	//start the server
	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//trap the signterm or interrupt and gracefully shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
