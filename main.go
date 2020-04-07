package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"muelltermine/api"
	"muelltermine/loader"
)

func printNow() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

func main() {
	printNow()
	fmt.Println("Hello Muelltermine!")

	addressLoader := loader.NewAddressLoader("data.json")

	addresses, err := addressLoader.LoadAddresses()

	if err != nil {
		panic(err)
	}

	printNow()

	router := mux.NewRouter()

	api.NewAddressesApi(addresses, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
