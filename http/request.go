// Package http holds all files related to the http requests made by stresser
package http

import (
	"bytes"
	fan "gist.github.com/titolins/4a0af7462eeb6bb76599e608a49d04b0.git"
	"net/http"
)

// RequestGen -> A wrapper whose only functionality is creating new
// http.Requests considering the reusability on those for post/put requests is
// questionable:
// https://github.com/golang/go/issues/19653
type RequestGen struct {
	Method string
	URI    string
	Data   *fan.Reader
}

// Map with accepted http methods so we can test for accepted values easily
var acceptedMethods = map[string]bool{
	"GET": false, "POST": true, "PUT": true, "DELETE": false, "PATCH": false}

// panics if we get an unexpected/ value
func panicIfNotMethodAccepted(method string) {
	if !acceptedMethods[method] {
		panic("Method not implemented or non-existent")
	}
}

// NewRequestGen -> builds and returns a http.Request wrapper (so we can
// generate those easily later when actually making requests)
func NewRequestGen(method string, uri string, data []byte) *RequestGen {
	// check if method accepted
	panicIfNotMethodAccepted(method)

	dataBuffer := fan.NewReader(bytes.NewBuffer(data))

	return &RequestGen{
		Method: method,
		URI:    uri,
		Data:   dataBuffer,
	}

}

// GenHTTPRequest -> Returns a new *http.Request to be used with the http.Client
func (req *RequestGen) GenHTTPRequest() *http.Request {
	httpReq, err := http.NewRequest(req.Method, req.URI, req.Data.View())
	if err != nil {
		panic(err.Error())
	}
	httpReq.Header.Add("Content-Type", "application/json")

	return httpReq
}
