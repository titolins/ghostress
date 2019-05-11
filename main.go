package main

import (
	"flag"
	"fmt"
	"github.com/titolins/ghostress/http"
	"github.com/titolins/ghostress/json"
	"io/ioutil"
)

func main() {
	var (
		useFile  bool
		nReq     int
		timeout  int
		method   string
		uri      string
		dataFile string
		data     []byte
		err      error
	)

	flag.IntVar(&nReq, "n_req", 1, "number of requests to be made to server")
	flag.IntVar(&timeout, "timeout", 0, "timeout between requests (seconds)")
	flag.StringVar(&method, "method", "PUT", "method to be used (POST or PUT)")
	flag.StringVar(&uri, "uri", "http://localhost:3000", "request uri")
	flag.StringVar(
		&dataFile,
		"data",
		"",
		("json data or descriptor file (depending on the value of the" +
			"'use_data_file' parameter)"))
	flag.BoolVar(
		&useFile,
		"use_data_file",
		true,
		("boolean indicating if a data file should be used (if false, a json" +
			" descriptor should be passed as '-data')"))
	flag.Parse()

	if dataFile == "" {
		panic("parameter 'data' cannot be null")
	}

	if useFile {
		// reads json file
		data, err = ioutil.ReadFile(dataFile)
		if err != nil {
			panic("Failed to read json file")
		}
	} else {
		descriptor := json.NewDescriptor(dataFile)
		fmt.Printf("descriptor = %+v\n", descriptor)
		generator := &json.Generator{Descriptor: descriptor}
		data = generator.GetData()
		fmt.Printf("generator.BuildObject() = %+v\n", string(data))
	}
	req := http.NewRequest(method, uri, data)

	stresser := &http.Stresser{
		Request: req,
		NReq:    nReq,
		Timeout: timeout}

	stresser.Stress()
}
