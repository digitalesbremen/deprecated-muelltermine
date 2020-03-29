package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type AddressLoader struct {
	FileName string
}

type JSONTime struct {
	time.Time
}

type Addresses struct {
	Addresses []Address
}

type Address struct {
	Street          string
	HouseNumber     string
	CollectionDates []GarbageCollectionDate
}

type GarbageCollectionDate struct {
	Date  JSONTime
	Types []string
}

func NewAddressLoader(filename string) *AddressLoader {
	addressLoader := AddressLoader{
		FileName: filename,
	}
	return &addressLoader
}

func (l *AddressLoader) LoadAddresses() (content Addresses, err error) {
	fmt.Println("Read file", l.FileName)
	dat, err := ioutil.ReadFile(l.FileName)

	if err != nil {
		return Addresses{}, err
	}

	var addresses Addresses

	fmt.Println("Unmarshal addresses from file", l.FileName)

	err = json.Unmarshal(dat, &addresses)

	if err != nil {
		return Addresses{}, err
	}

	fmt.Println("Addresses loaded:", len(addresses.Addresses))

	return addresses, nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	parsedDate, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*t = JSONTime{parsedDate}
	return nil
}
