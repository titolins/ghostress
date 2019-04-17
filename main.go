package main

import (
	"flag"
)

func main() {
	var (
		nReq    int
		timeout int
		method  string
		uri     string
		data    string
	)

	flag.IntVar(&nReq, "n_req", 1, "number of requests to be made to server")
	flag.IntVar(&timeout, "timeout", 0, "timeout between requests (seconds)")
	flag.StringVar(&method, "method", "PUT", "method to be used (POST or PUT)")
	flag.StringVar(&uri, "uri", "http://localhost:3000", "request uri")
	flag.StringVar(
		&data, "data", "test_data.json", "json file with the data to be sent")

	flag.Parse()

	req := NewRequest(method, uri, data)

	stresser := &Stresser{
		Request: req,
		NReq:    nReq,
		Timeout: timeout}

	stresser.Stress()

}
