package loader

import (
	"testing"
)

func TestAddressLoader_LoadAddresses(t *testing.T) {
	addressLoader := NewAddressLoader("testdata/addresses.json")
	addresses, _ := addressLoader.LoadAddresses()

	verifyStreet(t, `Aachener Straße`, addresses[0].Street)
	verifyHouseNumber(t, `2`, addresses[0].HouseNumber)

	verifyStreet(t, `Waiblinger Weg`, addresses[1].Street)
	verifyHouseNumber(t, `33`, addresses[1].HouseNumber)
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
