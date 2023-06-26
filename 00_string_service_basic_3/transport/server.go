package transport

import (
	"context"
	"log"
	"net/http"
	"time"

	httpTrans "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(_ context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/uppercase").Handler(httpTrans.NewServer(endpoints.Upperstring, decodeUppercaseRequest, encodeResponse))
	r.Methods("POST").Path("/count").Handler(httpTrans.NewServer(endpoints.Countstring, decodeCountRequest, encodeResponse))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before request processing
		startTm := time.Now()
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		// after request processing
		log.Printf("%s - Total Processing Time %s\n", r.RequestURI, time.Since(startTm))
	})
}
