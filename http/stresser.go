package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Stresser -> Describes and runs the request batch
type Stresser struct {
	Request *Request
	NReq    int
	Timeout int
}

// req -> Makes a single request
func (stresser *Stresser) req(ch chan<- string) {
	var res string
	httpClient := &http.Client{}
	reqRes, err := httpClient.Do(stresser.Request.GetHTTPRequest())
	if err != nil {
		res = fmt.Sprintf("Error making request\n*********\n%s\n", err.Error())
	} else {
		defer reqRes.Body.Close()
		//
		rbody, rerr := ioutil.ReadAll(reqRes.Body)
		if rerr != nil {
			res = fmt.Sprintf(
				"Error reading response\n***********\n%s\n", rerr.Error())
		} else {
			res = fmt.Sprintf(
				"Request response\n*******************\n%s\n", rbody)
		}
	}
	ch <- res
}

// Stress -> Starts the batch request
func (stresser *Stresser) Stress() {
	ch := make(chan string)

	fmt.Println("Starting stress test")
	fmt.Printf("stresser.NReq = %+v\n", stresser.NReq)
	for i := 0; i < stresser.NReq; i++ {
		fmt.Printf("i = %+v\n", i)
		go stresser.req(ch)
		time.Sleep(time.Duration(stresser.Timeout) * time.Second)
	}

	for i := 0; i < stresser.NReq; i++ {
		fmt.Printf("<-ch = %+v\n", <-ch)
	}

	fmt.Println("Finished stress test")
}
