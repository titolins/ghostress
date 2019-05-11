package http

import (
	"net/http"
)

// RequestSummary -> holds a reference to both the response and request objects
type RequestSummary struct {
	Request  *http.Request
	Response *http.Response
	ReqErr   error
}

// StressResult -> holds the responses of the requests batch made
type StressResult struct {
	res []RequestSummary
}

// NewStressResult -> returns a StressResult with the initiated res slice
func NewStressResult(nReq int) *StressResult {
	return &StressResult{
		res: make([]RequestSummary, nReq),
	}
}

// SetResult -> set the result of a given request to the response array
func (stressRes *StressResult) SetResult(reqSum RequestSummary, n int) {
	stressRes.res[n] = reqSum
}
