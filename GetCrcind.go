package crcind

import (
	"bytes"
	"log"
	"net/http"

	xj "github.com/basgys/goxml2json"
)

func GetCrcind(query string) *bytes.Buffer {
	res, err := http.Get("http://www.crcind.com/csp/samples/SOAP.Demo.cls?soap_method=GetByName&name=" + query)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := xj.Convert(res.Body)

	return json
}