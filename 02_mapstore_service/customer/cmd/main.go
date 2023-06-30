package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"gokit-microservices/02_mapstore_service/customer"
	"gokit-microservices/02_mapstore_service/customer/implementation"
	"gokit-microservices/02_mapstore_service/customer/mapstoredb"
	"gokit-microservices/02_mapstore_service/customer/transport"
	httptransport "gokit-microservices/02_mapstore_service/customer/transport/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {

	//1. command line var like http port to listen
	var httpAddr = flag.String("http", ":8085", "http port num")
	flag.Parse()

	//2. logger setup
	var logger log.Logger
	{
		//2.1 get NewLogfmtLogger logger instance
		logger = log.NewLogfmtLogger(os.Stderr)
		//2.2 get log syncing concurrently feature
		logger = log.NewSyncLogger(logger)
		//2.3 get log contextual
		logger = log.With(logger,
			"service", "customer",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	//3. start logging with level.Info for service start and end
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	//4. db variable declaration to be passed in repository

	//5. create service instance
	var srvc customer.CustomerService
	{ //5.1 repository instance with logger input
		repository, err := mapstoredb.NewRepository(logger)
		if err != nil {
			level.Error(logger).Log("repo-err", err)
		}
		//5.2 service instance with logger & repository input
		srvc = implementation.NewCustomerService(repository, logger)

		//5.3 add middleware if needed
	}

	//6. create go-kit endpoints for our service methods
	var endpoints transport.Endpoints
	{
		//6.1 endpoint instance
		endpoints = transport.NewEndpoints(srvc)

	}

	//7. create header for our endpoints
	var h http.Handler
	h = httptransport.NewService(endpoints, logger)

	//8. create err channel for handler any interruption
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	//9. Assigning port and ListenAndServe
	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:        *httpAddr,
			Handler:     h,
			ReadTimeout: 30,
		}
		errs <- server.ListenAndServe()
	}()
	level.Error(logger).Log("exit", <-errs)
}
