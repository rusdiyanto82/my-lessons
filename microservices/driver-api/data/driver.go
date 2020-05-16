package data

import "time"

type Driver struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	SKU       string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

func GetDrivers() []*Driver {
	return driverList
}

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
