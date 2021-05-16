package net

import (
	"github.com/chiupc/wc-api-go/request"
	"net/http"
)

// Sender provides HTTP Requests
type Sender struct {
	requestEnricher RequestEnricher
	urlBuilder      URLBuilder
	httpClient      Client
	requestCreator  RequestCreator
}

// Send method sends requests to WooCommerce API
func (s *Sender) Send(req request.Request) (resp *http.Response, err error) {
	request := s.prepareRequest(req)
	return s.httpClient.Do(request)
}

func (s *Sender) prepareRequest(req request.Request) *http.Request {
	URL := s.urlBuilder.GetURL(req)
	var request *http.Request
	//testBody := "{\n  \"payment_method\": \"bacs\",\n  \"payment_method_title\": \"Direct Bank Transfer\",\n  \"set_paid\": true,\n  \"billing\": {\n    \"first_name\": \"Lulu\",\n    \"last_name\": \"Lu\",\n    \"address_1\": \"Restoran Kwai Sun\",\n    \"address_2\": \"No. 5G &7G, Jalan SS15/4G, Ss 15\",\n    \"city\": \"Subang Jaya\",\n    \"state\": \"Selangor\",\n    \"postcode\": \"47500\",\n    \"country\": \"MY\",\n    \"email\": \"lulu@gmail.com\",\n    \"phone\": \"0170170170\"\n  },\n  \"shipping\": {\n    \"first_name\": \"Lulu\",\n    \"last_name\": \"Lu\",\n    \"address_1\": \"Restoran Kwai Sun\",\n    \"address_2\": \"No. 5G &7G, Jalan SS15/4G, Ss 15\",\n    \"city\": \"Subang Jaya\",\n    \"state\": \"Selangor\",\n    \"postcode\": \"47500\",\n    \"country\": \"MY\"\n  },\n  \"line_items\": [\n    {\n      \"product_id\": 12,\n      \"quantity\": 1\n    }\n  ],\n  \"shipping_lines\": [\n    {\n      \"method_id\": \"flat_rate\",\n      \"method_title\": \"Flat Rate\",\n      \"total\": \"10.00\"\n    }\n  ]\n}"
	if req.Body != nil{
		request, _ = s.requestCreator.NewRequest(req.Method, URL, req.Body)
		request.Header.Add("Content-Type", "application/json")
	}else {
		request, _ = s.requestCreator.NewRequest(req.Method, URL, nil)
	}
	s.requestEnricher.EnrichRequest(request, URL)
	if req.Values != nil && ("POST" == req.Method || "PUT" == req.Method) {
		request.Form = req.Values
	}

	return request
}

// SetRequestEnricher ...
func (s *Sender) SetRequestEnricher(a RequestEnricher) {
	s.requestEnricher = a
}

// SetURLBuilder ...
func (s *Sender) SetURLBuilder(urlBuilder URLBuilder) {
	s.urlBuilder = urlBuilder
}

// SetHTTPClient ...
func (s *Sender) SetHTTPClient(c Client) {
	s.httpClient = c
}

// SetRequestCreator ...
func (s *Sender) SetRequestCreator(rc RequestCreator) {
	s.requestCreator = rc
}
