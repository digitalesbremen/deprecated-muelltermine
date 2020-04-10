package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type AddressLoader struct {
	fileName string
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
		fileName: filename,
	}
	return &addressLoader
}

func (l *AddressLoader) LoadAddresses() (content []Address, err error) {
	fmt.Println("Read file", l.fileName)
	dat, err := ioutil.ReadFile(l.fileName)

	if err != nil {
		return []Address{}, err
	}

	var addresses Addresses

	fmt.Println("Unmarshal addresses from file", l.fileName)

	err = json.Unmarshal(dat, &addresses)

	if err != nil {
		return []Address{}, err
	}

	fmt.Println("Addresses loaded:", len(addresses.Addresses))

	return addresses.Addresses, nil
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
