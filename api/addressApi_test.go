package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

var testAddresses = newAddressBuilder().
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

func TestAddressApi_LoadAddressesWithoutQueryParameter(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address")

	verifyResponseHeader(t, res)
	verifyAddressSize(t, dtos, 10)
	verifyAddresses(t, []struct {
		index int
		got   string
		want  string
	}{
		{0, "Langwedeler Straße 1", dtos.Addresses[0]},
		{1, "Langwedeler Straße 1a", dtos.Addresses[1]},
		{2, "Langwedeler Straße 1b", dtos.Addresses[2]},
		{3, "Langwedeler Straße 2", dtos.Addresses[3]},
		{4, "Langwedeler Straße 3", dtos.Addresses[4]},
		{5, "Langwedeler Straße 3a", dtos.Addresses[5]},
		{6, "Langwedeler Straße 5", dtos.Addresses[6]},
		{7, "Langwedeler Straße 5a", dtos.Addresses[7]},
		{8, "Langwedeler Straße 5b", dtos.Addresses[8]},
		{9, "Zwoller Straße 38", dtos.Addresses[9]},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameter(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address?search=zwolle")

	verifyResponseHeader(t, res)
	verifyAddressSize(t, dtos, 3)
	verifyAddresses(t, []struct {
		index int
		got   string
		want  string
	}{
		{0, "Zwoller Straße 38", dtos.Addresses[0]},
		{1, "Zwoller Straße 40", dtos.Addresses[1]},
		{2, "Zwoller Straße 42", dtos.Addresses[2]},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameterNotFound(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address?search=not-found")

	verifyResponseHeader(t, res)
	verifyAddressSize(t, dtos, 0)
}

func sendRequest(t *testing.T, url string) (*httptest.ResponseRecorder, AddressesDto) {
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	NewAddressesApi(testAddresses, router)
	router.ServeHTTP(res, req)

	dtos := unmarshalResponse(t, res)
	return res, dtos
}

func verifyAddresses(t *testing.T, tests []struct {
	index int
	got   string
	want  string
}) {
	for _, tt := range tests {
		name := "GET /api/address [" + strconv.Itoa(tt.index) + "]"
		t.Run(name, func(t *testing.T) {
			verifyAddress(t, tt.index, tt.got, tt.want)
		})
	}
}

func verifyAddressSize(t *testing.T, dtos AddressesDto, want int) bool {
	return t.Run("verify body addresses length", func(t *testing.T) {
		if len(dtos.Addresses) > 10 {
			t.Errorf(`GET /api/address length = %d ; want %d`, len(dtos.Addresses), want)
		}
	})
}

func unmarshalResponse(t *testing.T, res *httptest.ResponseRecorder) AddressesDto {
	var dtos AddressesDto

	err := json.Unmarshal(res.Body.Bytes(), &dtos)

	t.Run("verify body is json", func(t *testing.T) {
		if err != nil {
			t.Errorf(`GET /api/address is no json: %s`, err.Error())
		}
	})
	return dtos
}

func verifyResponseHeader(t *testing.T, res *httptest.ResponseRecorder) bool {
	return t.Run("verify response header", func(t *testing.T) {
		if res.Header().Get("Content-Type") != "application/json" {
			t.Errorf(`GET /api/address header = %s ; want %s`, res.Header().Get("Content-Type"), "application/json")
		}
	})
}

func verifyAddress(t *testing.T, index int, got string, want string) {
	if want != got {
		t.Errorf(`GET /api/address [%d] = %s ; want %s`, index, got, want)
	}
}
