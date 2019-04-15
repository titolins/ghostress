package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	var nReq int
	var timeout int
	var err error

	if nReq, err = strconv.Atoi(os.Args[1]); err != nil {
		panic("nReq should be an int")
	}
	if timeout, err = strconv.Atoi(os.Args[2]); err != nil {
		panic("timeout should be an int")
	}
	method := os.Args[3]
	uri := os.Args[4]
	dataFile := os.Args[5]
	req := NewRequest(method, uri, dataFile)

	stresser := &Stresser{
		Request: req,
		NReq:    nReq,
		Timeout: time.Duration(timeout)}

	stresser.Stress()

}
