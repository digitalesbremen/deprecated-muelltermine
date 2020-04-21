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

	verifyContentTypeHeader(t, res, "application/hal+json")
	verifyStatusCode(t, res, 200)
	verifyStreetsLength(t, dtos.Streets, 11, `GET /api/address length = %d ; want %d`)
	verifyContent(t, "/api/address", []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Streets[0].StreetName, "Am Querkamp"},
		{1, dtos.Streets[1].StreetName, "Langwedeler Straße"},
		{2, dtos.Streets[2].StreetName, "Riensberger Straße"},
		{3, dtos.Streets[3].StreetName, "Schaffenrathstraße"},
		{4, dtos.Streets[4].StreetName, "Schaffhauser Straße"},
		{5, dtos.Streets[5].StreetName, "Steffensweg"},
		{6, dtos.Streets[6].StreetName, "Turnerstraße"},
		{7, dtos.Streets[7].StreetName, "Twiedelftsweg"},
		{8, dtos.Streets[8].StreetName, "Voltastraße"},
		{9, dtos.Streets[9].StreetName, "Von-Line-Straße"},
		{9, dtos.Streets[10].StreetName, "Zwoller Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameter(t *testing.T) {
	res := sendRequest("/api/address?search=zwolle", testAddresses)
	dtos := unmarshalStreetsDtoResponse(t, res)

	verifyContentTypeHeader(t, res, "application/hal+json")
	verifyStatusCode(t, res, 200)
	verifyStreetsLength(t, dtos.Streets, 1, `GET /api/address length = %d ; want %d`)
	verifyContent(t, "/api/address", []struct {
		index int
		got   string
		want  string
	}{
		{0, dtos.Streets[0].StreetName, "Zwoller Straße"},
	})
}

func TestAddressApi_LoadAddressesWithQueryParameterNotFound(t *testing.T) {
	res := sendRequest("/api/address?search=not-found", testAddresses)
	dtos := unmarshalStreetsDtoResponse(t, res)

	verifyContentTypeHeader(t, res, "application/hal+json")
	verifyStatusCode(t, res, 200)
	verifyStreetsLength(t, dtos.Streets, 0, `GET /api/address length = %d ; want %d`)
}

func TestAddressApi_LoadHouseNumbers(t *testing.T) {
	res := sendRequest("/api/address/Langwedeler%20Straße", testAddresses)
	dtos := unmarshalHouseNumbersDtoResponse(t, res)

	verifyContentTypeHeader(t, res, "application/json")
	verifyStatusCode(t, res, 200)
	verifyHouseNumberLength(t, dtos.HouseNumber, 3, `GET /api/address/Langwedeler%20Straße length = %d ; want %d`)
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

	verifyContentTypeHeader(t, res, "text/html")
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

func verifyStreetsLength(t *testing.T, values []StreetDto, want int, errorMessage string) {
	t.Run("verify body streets length", func(t *testing.T) {
		if len(values) != want {
			t.Errorf(errorMessage, len(values), want)
		}
	})
}

func verifyHouseNumberLength(t *testing.T, values []string, want int, errorMessage string) {
	t.Run("verify body house number length", func(t *testing.T) {
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

func verifyContentTypeHeader(t *testing.T, res *httptest.ResponseRecorder, contentType string) {
	t.Run("verify response header", func(t *testing.T) {
		if res.Header().Get("Content-Type") != contentType {
			t.Errorf(`GET /api/address header = %s ; want %s`, res.Header().Get("Content-Type"), contentType)
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
