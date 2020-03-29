package loader

import (
	"testing"
)

func TestAddressLoader_LoadAddresses(t *testing.T) {
	addressLoader := NewAddressLoader("testdata/addresses.json")
	addresses, _ := addressLoader.LoadAddresses()

	address1 := addresses[0]
	address2 := addresses[1]

	verifyStreet(t, `Aachener Straße`, address1.Street)
	verifyHouseNumber(t, `2`, address1.HouseNumber)

	verifyStreet(t, `Waiblinger Weg`, address2.Street)
	verifyHouseNumber(t, `33`, address2.HouseNumber)
}

func verifyStreet(t *testing.T, want string, got string) {
	if want != got {
		t.Errorf(`LoadAddresses()[0].Street = %s; want %s`, got, want)
	}
}

func verifyHouseNumber(t *testing.T, want string, got string) {
	if want != got {
		t.Errorf(`LoadAddresses()[0].HouseNumber = %s; want %s`, got, want)
	}
}
