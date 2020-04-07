package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"muelltermine/loader"
)

type Addresses struct {
	Addresses []string `json:"addresses,nilasempty"`
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
		HandleFunc("/api/address", addressLoader.allAddressesHandler).
		Queries("search", "{search:.*}").
		//Headers("Content-Type", "application/json").
		Methods("GET")
	addressLoader.router.
		HandleFunc("/api/address", addressLoader.allAddressesHandler).
		//Headers("Content-Type", "application/json").
		Methods("GET")

	return &addressLoader
}

func (a *AddressesApi) allAddressesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchValue := vars["search"]

	var addresses Addresses

	for _, entry := range a.addresses {
		fullQualifiedAddress := entry.Street + " " + entry.HouseNumber

		if containsIgnoreCase(fullQualifiedAddress, searchValue) {
			addresses.Addresses = append(addresses.Addresses, fullQualifiedAddress)
		}

		if len(addresses.Addresses) >= 10 {
			break
		}
	}

	w.Header().Add("Content-Type", "application/json")

	if len(addresses.Addresses) == 0 {
		_, _ = fmt.Fprint(w, `{"addresses":[]}`)
	} else {
		_ = json.NewEncoder(w).Encode(addresses)
	}
}

func containsIgnoreCase(s string, substring string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substring))
}
