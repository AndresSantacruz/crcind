package crcind

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetCrcind(query string) []byte {
	res, err := http.Get("http://www.crcind.com/csp/samples/SOAP.Demo.cls?soap_method=GetByName&name=" + query)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return data
}