// Модуль SOAP-клиента, иммитирует передачу данных к серверу по SOAP

package main

import (
	"context"
	"encoding/xml"
	"log"

	"github.com/3ideas/soap"
)

type FooRequest struct {
	XMLName xml.Name `xml:"Name"`
	Doc     string
}

// FooResponse a simple response
type FooResponse struct {
	Bar string
}

func main() {
	client := soap.NewClient("http://127.0.0.1:8088/", nil)
	client.Log = func(msg string, keyString_ValueInterface ...interface{}) {
		keyString_ValueInterface = append(keyString_ValueInterface, msg)
		log.Println(keyString_ValueInterface...)
	} // verbose
	response := &FooResponse{}
	httpResponse, err := client.Call(context.Background(), "operationFoo", &FooRequest{Doc: "ISimpleDocuments"}, response)
	if err != nil {
		panic(err)
	}
	log.Println(response.Bar, httpResponse.Status)
}
