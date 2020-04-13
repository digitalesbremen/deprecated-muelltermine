package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"muelltermine/loader"
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
	res := sendRequest("/api/address", testAddresses)
	dtos := unmarshalStreetsDtoResponse(t, res)

	verifyResponseHeader(t, res)
	verifyStatusCode(t, res, 200)
	verifyLength(t, dtos.Street, 10, `GET /api/address length = %d ; want %d`)
	verifyContent(t, "/api/address", []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Street[0], "Am Querkamp"},
		{1, dtos.Street[1], "Langwedeler Straße"},
		{2, dtos.Street[2], "Riensberger Straße"},
		{3, dtos.Street[3], "Schaffenrathstraße"},
		{4, dtos.Street[4], "Schaffhauser Straße"},
		{5, dtos.Street[5], "Steffensweg"},
		{6, dtos.Street[6], "Turnerstraße"},
		{7, dtos.Street[7], "Twiedelftsweg"},
		{8, dtos.Street[8], "Voltastraße"},
		{9, dtos.Street[9], "Von-Line-Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameter(t *testing.T) {
	res := sendRequest("/api/address?search=zwolle", testAddresses)
	dtos := unmarshalStreetsDtoResponse(t, res)

	verifyResponseHeader(t, res)
	verifyStatusCode(t, res, 200)
	verifyLength(t, dtos.Street, 1, `GET /api/address length = %d ; want %d`)
	verifyContent(t, "/api/address", []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Street[0], "Zwoller Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameterNotFound(t *testing.T) {
	res := sendRequest("/api/address?search=not-found", testAddresses)
	dtos := unmarshalStreetsDtoResponse(t, res)

	verifyResponseHeader(t, res)
	verifyStatusCode(t, res, 200)
	verifyLength(t, dtos.Street, 0, `GET /api/address length = %d ; want %d`)
}

func TestAddressApi_LoadHouseNumbers(t *testing.T) {
	res := sendRequest("/api/address/Langwedeler%20Straße", testAddresses)
	dtos := unmarshalHouseNumbersDtoResponse(t, res)

	verifyResponseHeader(t, res)
	verifyStatusCode(t, res, 200)
	verifyLength(t, dtos.HouseNumber, 3, `GET /api/address/Langwedeler%20Straße length = %d ; want %d`)
	verifyContent(t, "/api/address/Langwedeler%20Straße", []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.HouseNumber[0], "1"},
		{1, dtos.HouseNumber[1], "1a"},
		{2, dtos.HouseNumber[2], "1b"},
	})
}

func TestAddressApi_LoadHouseNumbersWithStreetNotExists(t *testing.T) {
	res := sendRequest("/api/address/NOT%20EXISTING", testAddresses)

	verifyStatusCode(t, res, 404)
}

func sendRequest(url string, addresses []loader.Address) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()

	router := mux.NewRouter()
	NewAddressesApi(addresses, router)
	router.ServeHTTP(res, req)

	return res
}

func verifyContent(t *testing.T, url string, tests []struct {
	index int
	got   string
	want  string
}) {
	for _, tt := range tests {
		name := "GET " + url + " [" + strconv.Itoa(tt.index) + "]"
		t.Run(name, func(t *testing.T) {
			if tt.want != tt.got {
				t.Errorf(`GET %s [%d] = %s ; want %s`, url, tt.index, tt.got, tt.want)
			}
		})
	}
}

func verifyLength(t *testing.T, values []string, want int, errorMessage string) {
	t.Run("verify body addresses length", func(t *testing.T) {
		if len(values) != want {
			t.Errorf(errorMessage, len(values), want)
		}
	})
}

func unmarshalStreetsDtoResponse(t *testing.T, res *httptest.ResponseRecorder) StreetsDto {
	var dtos StreetsDto

	err := json.Unmarshal(res.Body.Bytes(), &dtos)

	t.Run("verify body is json", func(t *testing.T) {
		if err != nil {
			t.Errorf(`GET /api/address is no json: %s`, err.Error())
		}
	})
	return dtos
}

func unmarshalHouseNumbersDtoResponse(t *testing.T, res *httptest.ResponseRecorder) HouseNumbersDto {
	var dtos HouseNumbersDto

	err := json.Unmarshal(res.Body.Bytes(), &dtos)

	t.Run("verify body is json", func(t *testing.T) {
		if err != nil {
			t.Errorf(`GET /api/address is no json: %s`, err.Error())
		}
	})
	return dtos
}

func verifyResponseHeader(t *testing.T, res *httptest.ResponseRecorder) {
	t.Run("verify response header", func(t *testing.T) {
		if res.Header().Get("Content-Type") != "application/json" {
			t.Errorf(`GET /api/address header = %s ; want %s`, res.Header().Get("Content-Type"), "application/json")
		}
	})
}

func verifyStatusCode(t *testing.T, res *httptest.ResponseRecorder, want int) {
	t.Run("verify http status code", func(t *testing.T) {
		if res.Code != want {
			t.Errorf(`GET /api/address code = %d ; want %d`, res.Code, want)
		}
	})
}
