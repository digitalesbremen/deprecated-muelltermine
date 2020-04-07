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

	if dtos.Addresses[0] != "Langwedeler Straße 1" {
		t.Errorf(`GET /api/address [0] = %s ; want %s`, dtos.Addresses[0], "Langwedeler Straße 1")
	}
}
