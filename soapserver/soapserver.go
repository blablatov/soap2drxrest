// Модуль парсер SOAP-сервера, для работы с сервисом интеграции Directum RX, через вызов модуля POST-запроса
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"odatapost"
	"time"

	"github.com/3ideas/soap"
)

// FooRequest a simple request
type FooRequest struct {
	XMLName xml.Name `xml:"Name"`
	Doc     string
}

// FooResponse a simple response
type FooResponse struct {
	Bar string
}

func main() {
	RunServer()
}

// RunServer run a little demo server
func RunServer() {
	soapServer := soap.NewServer()
	soapServer.Log = func(msg string, keyString_ValueInterface ...interface{}) {
		keyString_ValueInterface = append(keyString_ValueInterface, msg)
		var values []interface{}
		values = append(values, "SOAP Server: ")
		values = append(values, keyString_ValueInterface...)
		log.Println(values...)
	}

	soapServer.RegisterHandler(
		"/",            // Path
		"operationFoo", // SOAPAction
		"Name",         // tagname of soap body content // тэг содержимого тела soap-запроса
		func() interface{} { // RequestFactoryFunc - returns struct to unmarshal the request into
			return &FooRequest{}
		},
		// OperationHandlerFunc - do something
		func(request interface{}, w http.ResponseWriter, httpRequest *http.Request) (response interface{}, err error) {
			start := time.Now()
			fooRequest := request.(*FooRequest)
			fooResponse := &FooResponse{
				Bar: "SOAP-сервер: вызов метода сервиса интеграции: " + fooRequest.Doc,
			}
			response = fooResponse
			fmt.Println("Response server: ", response)

			var podata string
			switch fooRequest.Doc {
			case "ISimpleDocuments":
				podata = "Простой документ"
				cp := make(chan string) // канал функции post-запросов
				go odatapost.PostCreate(podata, cp)
				log.Println("\n\nРезультат post-запроса: \n", <-cp) // Получение данных запроса из канала горутины
				secs := time.Since(start).Seconds()
				fmt.Printf("%.2fs время выполнения запроса\n", secs)
			case "IUsers":
				podata = "Пользователь"
			default:
				fmt.Println("Данные строки не найдены!")
			}

			return
		},
	)
	err := http.ListenAndServe(":8088", soapServer)
	fmt.Println("exiting with error", err)
}
