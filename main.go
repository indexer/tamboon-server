package main

import (
	"github.com/omise/omise-go"
	"net/http"
	"os"
)

func main() {
	skey := os.Getenv("OMISE_SKEY")
	pkey := os.Getenv("OMISE_PKEY")
	if skey == "" {
		panic("Please set OMISE_SKEY")
	}
	if pkey == "" {
		panic("Please set OMISE_PKEY")
	}

	client, e := omise.NewClient(pkey, skey)
	if e != nil {
		panic(e)
	}

	if e := http.ListenAndServe(":8080", &TamboonHandler{client}); e != nil {
		panic(e)
	}
}
