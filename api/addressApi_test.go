package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAddressApi_LoadAddresses(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/address", nil)
	res := httptest.NewRecorder()

	addresses := newAddressBuilder().
		withAddress("Langwedeler Straße", "1").
		withAddress("Langwedeler Straße", "1a").
		withAddress("Langwedeler Straße", "1b").
		withAddress("Langwedeler Straße", "2").
		withAddress("Langwedeler Straße", "3").
		withAddress("Langwedeler Straße", "3a").
		withAddress("Langwedeler Straße", "5").
		withAddress("Langwedeler Straße", "5a").
		withAddress("Langwedeler Straße", "5b").
		withAddress("Zwoller Straße", "38").
		withAddress("Zwoller Straße", "40").
		withAddress("Zwoller Straße", "42").
		build()

	router := mux.NewRouter()
	NewAddressesApi(addresses, router)
	router.ServeHTTP(res, req)

	var dtos AddressesDto

	err := json.Unmarshal(res.Body.Bytes(), &dtos)

	if err != nil {
		t.Errorf(`GET /api/address is no json: %s`, err.Error())
	}

	if len(dtos.Addresses) > 10 {
		t.Errorf(`GET /api/address length = %d ; want %d`, len(dtos.Addresses), 10)
	}

	verifyAddress(t, 0, "Langwedeler Straße 1", dtos.Addresses[0])
	verifyAddress(t, 1, "Langwedeler Straße 1a", dtos.Addresses[1])
	verifyAddress(t, 2, "Langwedeler Straße 1b", dtos.Addresses[2])
	verifyAddress(t, 3, "Langwedeler Straße 2", dtos.Addresses[3])
	verifyAddress(t, 4, "Langwedeler Straße 3", dtos.Addresses[4])
	verifyAddress(t, 5, "Langwedeler Straße 3a", dtos.Addresses[5])
	verifyAddress(t, 6, "Langwedeler Straße 5", dtos.Addresses[6])
	verifyAddress(t, 7, "Langwedeler Straße 5a", dtos.Addresses[7])
	verifyAddress(t, 8, "Langwedeler Straße 5b", dtos.Addresses[8])
	verifyAddress(t, 9, "Zwoller Straße 38", dtos.Addresses[9])
}

func verifyAddress(t *testing.T, index int, got string, want string) {
	if want != got {
		t.Errorf(`GET /api/address [%d] = %s ; want %s`, index, got, want)
	}
}
