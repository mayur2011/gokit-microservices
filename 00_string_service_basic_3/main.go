package main

import (
	"context"
	"gokit-microservices/00_string_service_basic_3/service"
	"gokit-microservices/00_string_service_basic_3/transport"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting The String Service (3) ..!")
	stringSrvc := service.NewStringService()

	endpoints := transport.NewEndpoints(stringSrvc)
	handler := transport.NewHTTPServer(context.Background(), endpoints)
	log.Fatal(http.ListenAndServe(":8082", handler))
}
