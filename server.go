package main

import (
	"encoding/json"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"net/http"
)

type TamboonHandler struct {
	client *omise.Client
}

func (handler *TamboonHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	method, path := req.Method, req.URL.Path
	fmt.Printf("%s %s\n", method, path)

	if method == "GET" && path == "/" {
		handler.GET(resp, req)

	} else if method == "POST" && path == "/donate" {
		handler.POST(resp, req)

	} else {
		http.NotFound(resp, req)

	}
}

func (handler *TamboonHandler) GET(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type","application/json")//Set Content-Type Header
	if e := json.NewEncoder(resp).Encode(charities); e != nil {
		http.Error(resp, e.Error(), 400)
		return
	}
}

func (handler *TamboonHandler) POST(resp http.ResponseWriter, req *http.Request) {
	donation := &Donation{}
	defer req.Body.Close()
	resp.Header().Set("Content-Type","application/json")

	if e := json.NewDecoder(req.Body).Decode(donation); e != nil {
		resp.WriteHeader(http.StatusTeapot)
		json.NewEncoder(resp).Encode(&Error{Message: e.Error()})
		return
	}

	charge, operation := &omise.Charge{}, &operations.CreateCharge{
		Card:        donation.Token,
		Amount:      donation.Amount,
		Currency:    "THB",
		Description: donation.Name,
	}

	if e := handler.client.Do(charge, operation); e != nil {
		http.Error(resp, e.Error(), 400)
		return
	}

	if e := json.NewEncoder(resp).Encode(&Result{Success: true}); e != nil {
		http.Error(resp, e.Error(), 400)
		return
	}
}
