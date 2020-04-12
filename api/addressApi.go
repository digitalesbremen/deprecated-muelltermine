package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"muelltermine/loader"
)

type StreetsDto struct {
	Street []string `json:"streets"`
}

type HouseNumbersDto struct {
	HouseNumber []string `json:"houseNumbers"`
}

type AddressesApi struct {
	addresses []loader.Address
	router    *mux.Router
}

func NewAddressesApi(addresses []loader.Address, router *mux.Router) *AddressesApi {
	addressLoader := AddressesApi{
		addresses: addresses,
		router:    router,
	}

	addressLoader.router.
		HandleFunc("/api/address", addressLoader.getAllAddressesHandler).
		Queries("search", "{search:.*}").
		//Headers("Content-Type", "application/json").
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address", addressLoader.getAllAddressesHandler).
		//Headers("Content-Type", "application/json").
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address/{street}", addressLoader.getAddressHandler).
		//Headers("Content-Type", "application/json").
		Methods("GET")

	return &addressLoader
}

func (a *AddressesApi) getAllAddressesHandler(w http.ResponseWriter, r *http.Request) {
	searchValue := mux.Vars(r)["search"]

	var addresses StreetsDto

	for _, entry := range a.addresses {
		if containsIgnoreCase(entry.Street, searchValue) && !contains(addresses.Street, entry.Street) {
			addresses.Street = append(addresses.Street, entry.Street)
		}

		if len(addresses.Street) >= 10 {
			break
		}
	}

	w.Header().Add("Content-Type", "application/json")

	if len(addresses.Street) == 0 {
		_, _ = fmt.Fprint(w, `{"addresses":[]}`)
	} else {
		_ = json.NewEncoder(w).Encode(addresses)
	}
}

func (a *AddressesApi) getAddressHandler(w http.ResponseWriter, r *http.Request) {
	street := mux.Vars(r)["street"]

	var houseNumbers HouseNumbersDto

	for _, entry := range a.addresses {
		if strings.ToLower(entry.Street) == strings.ToLower(street) {
			houseNumbers.HouseNumber = append(houseNumbers.HouseNumber, entry.HouseNumber)
		}
	}

	_ = json.NewEncoder(w).Encode(houseNumbers)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsIgnoreCase(s string, substring string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substring))
}
