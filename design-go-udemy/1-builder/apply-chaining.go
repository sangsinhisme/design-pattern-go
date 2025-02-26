package main

import (
	"fmt"
	"net/url"
)

type Request struct {
	Method, URL, Body string
	Header            map[string]string
	QueryParams       url.Values
}

type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{&Request{
		Header:      make(map[string]string),
		QueryParams: make(url.Values),
	}}
}

func (rb *RequestBuilder) Build() *Request {
	return rb.request
}

func (rb *RequestBuilder) Method(method string) *RequestBuilder {
	rb.request.Method = method
	return rb
}

func (rb *RequestBuilder) URL(url string) *RequestBuilder {
	rb.request.URL = url
	return rb
}

func (rb *RequestBuilder) Body(body string) *RequestBuilder {
	rb.request.Body = body
	return rb
}

func (rb *RequestBuilder) Header(key, value string) *RequestBuilder {
	rb.request.Header[key] = value
	return rb
}

func (rb *RequestBuilder) QueryParam(key, value string) *RequestBuilder {
	rb.request.QueryParams.Add(key, value)
	return rb
}

func main() {
	request := NewRequestBuilder().
		Method("POST").
		URL("https://api.example.com").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token").
		QueryParam("user", "john").
		QueryParam("lang", "en").
		Body(`{"message": "hello"}`).
		Build()

	// Format query params for better visualization
	queryString := request.QueryParams.Encode()
	fullURL := request.URL
	if queryString != "" {
		fullURL += "?" + queryString
	}

	fmt.Println("Method:", request.Method)
	fmt.Println("URL:", fullURL)
	fmt.Println("Headers:", request.Header)
	fmt.Println("Query Params:", request.QueryParams.Encode())
	fmt.Println("Body:", request.Body)
}
