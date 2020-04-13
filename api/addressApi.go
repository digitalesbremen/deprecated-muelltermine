package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"muelltermine/api/page"
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
		HandleFunc("/api/address", addressLoader.getStreetsHandler).
		Queries("search", "{search:.*}").
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address/", addressLoader.getStreetsHandler).
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address", addressLoader.getStreetsHandler).
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address/{street}", addressLoader.getHouseNumbersHandler).
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address/{street}/", addressLoader.getHouseNumbersHandler).
		Methods("GET")

	return &addressLoader
}

func (a *AddressesApi) getStreetsHandler(w http.ResponseWriter, r *http.Request) {
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

func (a *AddressesApi) getHouseNumbersHandler(w http.ResponseWriter, r *http.Request) {
	street := mux.Vars(r)["street"]

	var houseNumbers HouseNumbersDto

	for _, entry := range a.addresses {
		if strings.ToLower(entry.Street) == strings.ToLower(street) {
			houseNumbers.HouseNumber = append(houseNumbers.HouseNumber, entry.HouseNumber)
		}
	}

	if len(houseNumbers.HouseNumber) == 0 {
		w.WriteHeader(http.StatusNotFound)
		page.Write404(w, r.Method+" "+r.URL.Path)
	} else {
		w.Header().Add("Content-Type", "application/json")

		_ = json.NewEncoder(w).Encode(houseNumbers)
	}
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
