package http

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

const summaryTemplate = `
Stress Result:
==============
Total requests made   : {{ .NReq }}
Total requests failed : {{ .NReqFailed }}
Success Rate          : {{ .GetRequestsSuccessRate }}%
`

// RequestSummary -> holds a reference to both the response and request objects
type RequestSummary struct {
	Request  *http.Request
	Response *http.Response
	ReqErr   error
}

// StressResult -> holds the responses of the requests batch made
type StressResult struct {
	res        []RequestSummary
	template   *template.Template
	NReq       int
	NReqFailed int
}

// NewStressResult -> returns a StressResult with the initiated res slice
func NewStressResult(nReq int) *StressResult {
	t, err := template.New("RequestSummary").Parse(summaryTemplate)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse summary template: %s", err))
	}

	return &StressResult{
		res:        make([]RequestSummary, nReq),
		template:   t,
		NReq:       nReq,
		NReqFailed: 0,
	}
}

// SetResult -> set the result of a given request to the response array
func (stressRes *StressResult) SetResult(reqSum RequestSummary, n int) {
	stressRes.res[n] = reqSum
	// increment the failed req counter (so we don't need to calculate it later)
	if reqSum.ReqErr != nil {
		stressRes.NReqFailed++
	}
}

// GetRequestsSuccessRate -> simply calculates the percentage using NReq and
// NReqFailed
func (stressRes *StressResult) GetRequestsSuccessRate() int {
	return (100 - (stressRes.NReqFailed/stressRes.NReq)*100)
}

// PrintSummary -> Prints the stress result summary
func (stressRes *StressResult) PrintSummary() {
	stressRes.template.Execute(os.Stdout, stressRes)
}
