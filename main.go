package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"muelltermine/loader"
)

func printNow() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

func main() {
	printNow()
	fmt.Println("Hello Golang!")

	addressLoader := loader.NewAddressLoader("data.json")

	_, err := addressLoader.LoadAddresses()

	if err != nil {
		panic(err)
	}

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
