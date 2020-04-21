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
	Streets []StreetDto `json:"streets"`
	Links   LinksDto    `json:"_links,omitempty"`
}

type StreetDto struct {
	StreetName string   `json:"streetName"`
	Links      LinksDto `json:"_links,omitempty"`
}

type LinksDto struct {
	Self LinkDto `json:"self,omitempty"`
}

type LinkDto struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
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
	requestUrl := r.URL.Path

	streetsDto := StreetsDto{
		Links: LinksDto{
			Self: LinkDto{
				Href:      requestUrl + "{?search}",
				Templated: true,
			},
		},
	}

	for _, entry := range a.addresses {
		var streetNames []string
		for _, street := range streetsDto.Streets {
			streetNames = append(streetNames, street.StreetName)
		}

		if containsIgnoreCase(entry.Street, searchValue) && !contains(streetNames, entry.Street) {
			streetsDto.Streets = append(streetsDto.Streets, StreetDto{
				StreetName: entry.Street,
				Links: LinksDto{
					Self: LinkDto{
						Href: requestUrl + "/" + entry.Street,
					},
				},
			})
		}
	}

	w.Header().Add("Content-Type", "application/json")

	if len(streetsDto.Streets) == 0 {
		_, _ = fmt.Fprint(w, `{"streets":[]}`)
	} else {
		_ = json.NewEncoder(w).Encode(streetsDto)
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
		w.Header().Add("Content-Type", "text/html")
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
