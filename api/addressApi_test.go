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
	withAddress("Am Querkamp", "36").
	withAddress("Langwedeler Straße", "1").
	withAddress("Langwedeler Straße", "1a").
	withAddress("Langwedeler Straße", "1b").
	withAddress("Riensberger Straße", "94").
	withAddress("Schaffenrathstraße", "44").
	withAddress("Schaffhauser Straße", "5").
	withAddress("Steffensweg", "78").
	withAddress("Turnerstraße", "295").
	withAddress("Twiedelftsweg", "5").
	withAddress("Voltastraße", "120").
	withAddress("Von-Line-Straße", "4").
	withAddress("Zwoller Straße", "38").
	withAddress("Zwoller Straße", "40").
	withAddress("Zwoller Straße", "42").
	build()

func TestAddressApi_LoadAddressesWithoutQueryParameter(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address")

	verifyResponseHeader(t, res)
	verifyAddressesLength(t, dtos, 2)
	verifyAddresses(t, []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Addresses[0], "Am Querkamp"},
		{1, dtos.Addresses[1], "Langwedeler Straße"},
		{2, dtos.Addresses[2], "Riensberger Straße"},
		{3, dtos.Addresses[3], "Schaffenrathstraße"},
		{4, dtos.Addresses[4], "Schaffhauser Straße"},
		{5, dtos.Addresses[5], "Steffensweg"},
		{6, dtos.Addresses[6], "Turnerstraße"},
		{7, dtos.Addresses[7], "Twiedelftsweg"},
		{8, dtos.Addresses[8], "Voltastraße"},
		{9, dtos.Addresses[9], "Von-Line-Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameter(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address?search=zwolle")

	verifyResponseHeader(t, res)
	verifyAddressesLength(t, dtos, 1)
	verifyAddresses(t, []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Addresses[0], "Zwoller Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameterNotFound(t *testing.T) {
	res, dtos := sendRequest(t, "/api/address?search=not-found")

	verifyResponseHeader(t, res)
	verifyAddressesLength(t, dtos, 0)
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

func verifyAddressesLength(t *testing.T, dtos AddressesDto, want int) bool {
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
