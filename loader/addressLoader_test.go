package loader

import (
	"testing"
)

func TestAddressLoader_LoadAddresses(t *testing.T) {
	addressLoader := NewAddressLoader("testdata/addresses.json")
	addresses, _ := addressLoader.LoadAddresses()

	address1 := addresses.Addresses[0]
	address2 := addresses.Addresses[1]

	if address1.Street != `Aachener Straße` {
		t.Errorf(`LoadAddresses()[0].Street = %s; want %s`, address1.Street, `Aachener Straße`)
	}

	if address2.Street != `Waiblinger Weg` {
		t.Errorf(`LoadAddresses()[0].Street = %s; want %s`, address1.Street, `Waiblinger Weg`)
	}
}
