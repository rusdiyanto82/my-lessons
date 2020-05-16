package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("running hello handler")

	//read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("error reading body", err)
		http.Error(rw, "unable to read request body", http.StatusBadRequest)
		return
	}

	//write the response
	fmt.Fprintf(rw, "hello %s", b)
}
