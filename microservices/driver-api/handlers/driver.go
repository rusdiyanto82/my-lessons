package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rusdiyanto82/my-lessons/microservices/driver-api/data"
)

type Driver struct {
	l *log.Logger
}

func NewDriver(l *log.Logger) *Driver {
	return &Driver{l}
}

func (p *Driver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	dr := data.GetDrivers()
	d, err := json.Marshal(dr)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}
