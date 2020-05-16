package handlers

import (
	"log"
	"net/http"

	"github.com/rusdiyanto82/my-lessons/microservices/driver-api/data"
)

//Drivers is a http.handler
type Drivers struct {
	l *log.Logger
}

//NewDriver creates a driver handler with a given logger
func NewDriver(l *log.Logger) *Drivers {
	return &Drivers{l}
}

//ServeHTTP is the main entry point for the handler and satifies the http.Handler interface
func (d *Drivers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		d.getDrivers(rw, r)
		return
	}

	//catch all
	//if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

//getDrivers return the drivers from the data
func (d *Drivers) getDrivers(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Drivers")

	//fetch the drivers from the data
	dr := data.GetDrivers()

	//serialize the list to Json
	err := dr.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
