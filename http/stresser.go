package http

import (
	"fmt"
	"net/http"
	"time"
)

// Stresser -> Describes and runs the request batch
type Stresser struct {
	RequestGen *RequestGen
	NReq       int
	Timeout    int
	Result     *StressResult
}

// NewStresser -> builds a stresser with a response object
func NewStresser(req *RequestGen, nReq int, timeout int) *Stresser {
	StressResult := NewStressResult(nReq)
	return &Stresser{
		RequestGen: req,
		NReq:       nReq,
		Timeout:    timeout,
		Result:     StressResult,
	}
}

// req -> Makes a single request
func (stresser *Stresser) req(resCh chan<- RequestSummary) {
	httpClient := &http.Client{}
	httpReq := stresser.RequestGen.GenHTTPRequest()
	start := time.Now()
	httpRes, err := httpClient.Do(httpReq)
	elapsed := time.Since(start).Seconds()
	summary := &RequestSummary{
		Request:     httpReq,
		Response:    httpRes,
		ReqErr:      err,
		TimeElapsed: elapsed,
	}
	resCh <- *summary
}

// Stress -> Starts the batch request
func (stresser *Stresser) Stress() {
	resCh := make(chan RequestSummary)

	fmt.Println("Starting stress test")
	fmt.Printf("stresser.NReq = %+v\n", stresser.NReq)
	for i := 0; i < stresser.NReq; i++ {
		fmt.Printf("i = %+v\n", i)
		go stresser.req(resCh)
		time.Sleep(time.Duration(stresser.Timeout) * time.Second)
	}

	for i := 0; i < stresser.NReq; i++ {
		summary := <-resCh
		fmt.Printf("<-ch = %+v\n", summary)
		stresser.Result.SetResult(summary, i)
	}

	fmt.Println("Finished stress test")
	stresser.Result.PrintSummary()
}
