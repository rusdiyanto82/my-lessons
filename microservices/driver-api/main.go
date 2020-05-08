package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//request to the path /goodbye will be handled by this function
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("running goodbye handler")
		fmt.Fprintf(rw, "goodbye")
	})

	//any other request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("running hello handler")

		//read the body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading body", err)
			http.Error(rw, "unable to read request body", http.StatusBadRequest)
			return
		}

		//write the response
		fmt.Fprintf(rw, "hello %s", b)
	})

	//listen for connection on all ip address (0.0.0.0)
	//port 9090
	log.Println("starting server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
