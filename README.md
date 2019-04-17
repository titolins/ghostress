# Ghostress

## Description
Simple golang application to perform multiple concurrent POST/PUT HTTP requests sending a json file as payload.

## Usage
The following flags are defined:
* `n_req`   -> number of requests to be made to server (defaults to 1)
* `timeout` -> the timeout (in seconds) between requests (defaults to 0)
* `method`  -> method used to send the payload - PUT or POST (defaults to `PUT`)
* `uri`     -> request address (defaults to `http://localhost:3000/`)
* `data`    -> path to data file to be sent as payload (defaults to `test_data.json`)

* Example:
```bash
$ # presuming you have a json payload file named `test_data.json` in the same
$ # folder, this will send a single PUT request to the uri defined
$ ./ghostress -uri=http://11.22.33.44:80/test/api
```









Author: Tito Lins
