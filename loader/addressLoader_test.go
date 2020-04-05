package loader

import (
	"testing"
	"time"
)

func TestAddressLoader_LoadAddresses(t *testing.T) {
	addressLoader := NewAddressLoader("testdata/addresses.json")
	addresses, _ := addressLoader.LoadAddresses()

	verifyStreet(t, 0, `Aachener Straße`, addresses[0].Street)
	verifyHouseNumber(t, 0, `2`, addresses[0].HouseNumber)
	verifyCollectionDatesDate(t, 0, 0, "2020-03-09", addresses[0].CollectionDates[0].Date.Time)
	verifyCollectionDatesDate(t, 0, 1, "2020-03-16", addresses[0].CollectionDates[1].Date.Time)

	verifyStreet(t, 1, `Waiblinger Weg`, addresses[1].Street)
	verifyHouseNumber(t, 1, `33`, addresses[1].HouseNumber)
	verifyCollectionDatesDate(t, 1, 0, "2020-03-11", addresses[1].CollectionDates[0].Date.Time)
}

func verifyStreet(t *testing.T, addressIndex int, want string, got string) {
	if want != got {
		t.Errorf(`LoadAddresses()[%d].Street = %s; want %s`, addressIndex, got, want)
	}
}

func verifyHouseNumber(t *testing.T, addressIndex int, want string, got string) {
	if want != got {
		t.Errorf(`LoadAddresses()[%d].HouseNumber = %s; want %s`, addressIndex, got, want)
	}
}

func verifyCollectionDatesDate(t *testing.T, addressIndex int, collectionDatesIndex int, want string, got time.Time) {
	if want != got.Format("2006-01-02") {
		t.Errorf(`LoadAddresses()[%d].CollectionDates[%d].Date = %s; want %s`, addressIndex, collectionDatesIndex, got, want)
	}
}
