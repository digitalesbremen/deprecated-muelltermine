package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func printNow() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

func main() {
	printNow()
	fmt.Println("Hello Golang!")

	dat, err := ioutil.ReadFile("data.json")

	if err != nil {
		panic(err)
	}

	var addresses Addresses

	_ = json.Unmarshal(dat, &addresses)

	fmt.Println("Addresses loaded: ", len(addresses.Addresses))
	printNow()

	router := mux.NewRouter()
	router.HandleFunc("/api/address", ExampleHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	parsedDate, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*t = JSONTime{parsedDate}
	return nil
}

type JSONTime struct {
	time.Time
}

type Addresses struct {
	Addresses []Address
}

type Address struct {
	Street          string
	HouseNumber     string
	CollectionDates []GarbageCollectionDate
}

type GarbageCollectionDate struct {
	Date  JSONTime
	Types []string
}
