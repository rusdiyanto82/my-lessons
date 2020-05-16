package data

import (
	"encoding/json"
	"io"
	"time"
)

//Driver defines the structure for an API driver
type Driver struct {
	ID        int    `json:"id"`
	NIP       string `json:"nip"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	SKU       string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

//Drivers is an collection of Driver
type Drivers []*Driver

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (d *Drivers) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

// GetDrivers returns a list of drivers
func GetDrivers() Drivers {
	return driverList
}

// driverList is a hard coded list of drivers for this
// example data source
var driverList = []*Driver{
	&Driver{
		ID:        1,
		NIP:       "00006962",
		Name:      "Rusdi",
		Address:   "Jl Mampang",
		SKU:       "123abc",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Driver{
		ID:        2,
		NIP:       "00006963",
		Name:      "Rusdiyanto",
		Address:   "Jl Mampang Prapatan",
		SKU:       "123abcd",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
