package loader

import (
	"testing"
)

func TestAddressLoader_LoadAddresses(t *testing.T) {
	addressLoader := NewAddressLoader("testdata/addresses.json")
	addresses, _ := addressLoader.LoadAddresses()

	verifyStreet(t, 0, `Aachener Straße`, addresses[0].Street)
	verifyHouseNumber(t, 0, `2`, addresses[0].HouseNumber)
	verifyCollectionDatesDate(t, 0, 0, "2020-03-09", []string{"BLUE", "YELLOW"}, addresses[0].CollectionDates[0])
	verifyCollectionDatesDate(t, 0, 1, "2020-03-16", []string{"BLACK", "BROWN"}, addresses[0].CollectionDates[1])

	verifyStreet(t, 1, `Waiblinger Weg`, addresses[1].Street)
	verifyHouseNumber(t, 1, `33`, addresses[1].HouseNumber)
	verifyCollectionDatesDate(t, 1, 0, "2020-03-11", []string{"BLUE", "YELLOW"}, addresses[1].CollectionDates[0])
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

func verifyCollectionDatesDate(t *testing.T, addressIndex int, collectionDatesIndex int, wantDate string, wantTypes []string, got GarbageCollectionDate) {
	formattedGotDate := got.Date.Format("2006-01-02")
	if wantDate != formattedGotDate {
		t.Errorf(`LoadAddresses()[%d].CollectionDates[%d].Date = %s; want %s`, addressIndex, collectionDatesIndex, formattedGotDate, wantDate)
	}

	if !equal(wantTypes, got.Types) {
		t.Errorf(`LoadAddresses()[%d].CollectionDates[%d].Types = %s; want %s`, addressIndex, collectionDatesIndex, got.Types, wantTypes)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
