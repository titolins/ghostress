package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

// overall result template
const stressResultTemplate = `
Stress Result:
==============

Total requests made   : {{ .NReq }}
Total requests failed : {{ .NReqFailed }}
Success Rate          : {{ .GetRequestsSuccessRate }}%
Average time elapsed  : {{ .AvgTime }}ns

==============
`

// individual responses template
const responseSummaryTemplate = `
Response n. {{ .ID }}

==============

Body        : {{ .BodyText }}
Status Code : {{ .StatusCode }}
Time elapsed: {{ .TimeElapsed }}ns

==============
`

type responseSummary struct {
	ID          int
	BodyText    string
	StatusCode  int
	TimeElapsed float64
	template    *template.Template
}

func newResponseSummary(requestSummary *RequestSummary, id int) *responseSummary {
	defer requestSummary.Response.Body.Close()
	body, err := ioutil.ReadAll(requestSummary.Response.Body)
	if err != nil {
		body = []byte(fmt.Sprintf("Error reading response body: %s\n", err))
	}

	t, err := template.New("ResponseSummary").Parse(responseSummaryTemplate)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse response summary template: %s", err))
	}

	summary := &responseSummary{
		ID:          id,
		BodyText:    string(body),
		StatusCode:  requestSummary.Response.StatusCode,
		TimeElapsed: requestSummary.TimeElapsed,
		template:    t,
	}

	return summary
}

// RequestSummary -> holds a reference to both the response and request objects
type RequestSummary struct {
	Request     *http.Request
	Response    *http.Response
	ReqErr      error
	TimeElapsed float64
}

// StressResult -> holds the responses of the requests batch made
type StressResult struct {
	res        []RequestSummary
	template   *template.Template
	NReq       int
	NReqFailed int
	AvgTime    float64
}

// NewStressResult -> returns a StressResult with the initiated res slice
func NewStressResult(nReq int) *StressResult {
	t, err := template.New("StressResult").Parse(stressResultTemplate)
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
	var totalTimeElapsed float64
	fmt.Println("Printing responses...")
	for i, summary := range stressRes.res {
		sumObj := newResponseSummary(&summary, i)
		sumObj.template.Execute(os.Stdout, sumObj)
		totalTimeElapsed += sumObj.TimeElapsed
	}
	stressRes.AvgTime = (totalTimeElapsed / float64(len(stressRes.res)))
	stressRes.template.Execute(os.Stdout, stressRes)
}
