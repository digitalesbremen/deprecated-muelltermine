package api

import "muelltermine/loader"

type AddressesBuilder struct {
	addresses []loader.Address
}

func newAddressBuilder() AddressesBuilder {
	return AddressesBuilder{addresses: []loader.Address{}}
}

func (b AddressesBuilder) withAddress(street string, houseNumber string) AddressesBuilder {
	address := loader.Address{
		Street:          street,
		HouseNumber:     houseNumber,
		CollectionDates: nil,
	}

	b.addresses = append(b.addresses, address)
	return b
}

func (b AddressesBuilder) build() []loader.Address {
	return b.addresses
}
