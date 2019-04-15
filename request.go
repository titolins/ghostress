// Main package for ghostresser
package main

import (
	"bytes"
	fan "gist.github.com/jmackie/11570bdcd8a4c10d72619a5e1f21c5f8.git"
	"io/ioutil"
	"net/http"
)

// Request -> A wrapper whose only functionality is creating new http.Requests
// considering the reusability on those for post/put requests is questionable:
// https://github.com/golang/go/issues/19653
type Request struct {
	Method string
	URI    string
	Data   *fan.Reader
}

// Map with accepted http methods so we can test for accepted values easily
var acceptedMethods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "DELETE": false, "PATCH": false}

// simply panics if we get an unexpected value
func panicIfNotMethodAccepted(method string) {
	if !acceptedMethods[method] {
		panic("Method not implemented or non-existent")
	}
}

// NewRequest -> builds and returns a http.Request with the right data
func NewRequest(method string, uri string, dataFile string) *Request {
	// first check if method accepted
	panicIfNotMethodAccepted(method)

	// then try to read the data file, and panics in case of error
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		panic("Failed to read json file")
	}

	// build the buffer from the []byte
	dataBuffer := fan.NewReader(bytes.NewBuffer(data))

	return &Request{
		Method: method,
		URI:    uri,
		Data:   dataBuffer,
	}

}

// GetHTTPRequest -> Returns a new *http.Request to be used with the http.Client
func (req *Request) GetHTTPRequest() *http.Request {
	// make the request
	httpReq, err := http.NewRequest(req.Method, req.URI, req.Data.View())
	if err != nil {
		panic(err.Error())
	}

	// set the headers
	//httpReq.ContentLength = int64(dataBuffer.Len())
	httpReq.Header.Add("Content-Type", "application/json")

	return httpReq
}
