package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"muelltermine/loader"
)

func TestAddressApi_LoadAddresses(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/address", nil)
	res := httptest.NewRecorder()

	var addresses []loader.Address
	addresses = append(addresses, buildAddress("Langwedeler Str.", "1"))

	router := mux.NewRouter()
	NewAddressesApi(addresses, router)
	router.ServeHTTP(res, req)

	var dtos Addresses

	err := json.Unmarshal(res.Body.Bytes(), &dtos)

	if err != nil {
		t.Errorf(`GET /api/address is no json: %s`, err.Error())
	}

	if dtos.Addresses[0] != "Langwedeler Str. 1" {
		t.Errorf(`GET /api/address [0] = %s ; want %s`, dtos.Addresses[0], "Langwedeler Str. 1")
	}
}

func buildAddress(street string, houseNumber string) loader.Address {
	return loader.Address{
		Street:          street,
		HouseNumber:     houseNumber,
		CollectionDates: nil,
	}
}
