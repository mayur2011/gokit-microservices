package main

import (
	"gokit-microservices/00_string_service_basic_2/service"
	trans "gokit-microservices/00_string_service_basic_2/transport"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	log.Println("Starting The String Service ..!")
	stringService := service.NewStringService()
	uppercaseStringHandler := httptransport.NewServer(trans.MakeUppercaseEndpoint(stringService), trans.DecodeUppercaseRequest, trans.EncodeResponse)
	countHandler := httptransport.NewServer(trans.MakeCountEndpoint(stringService), trans.DecodeCountRequest, trans.EncodeResponse)

	http.Handle("/uppercase", uppercaseStringHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
