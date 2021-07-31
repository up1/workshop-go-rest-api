package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afex/hystrix-go/hystrix"
)

const commandName = "producer_api"

func main() {

	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  50,
		RequestVolumeThreshold: 3,
		SleepWindow:            1000,
	})

	http.HandleFunc("/", logger(handle))
	log.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	output := make(chan bool, 1)
	errors := hystrix.Go(commandName, func() error {
		// talk to other services
		err := callChargeProducerAPI()
		// err := callWithRetryV1()

		if err == nil {
			output <- true
		}
		return err
	}, nil)

	select {
	case out := <-output:
		// success
		log.Printf("success %v", out)
	case err := <-errors:
		// failure
		log.Printf("failed %s", err)
	}
}

// logger is Handler wrapper function for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}

func callChargeProducerAPI() error {
	fmt.Println(os.Getenv("SERVER_ERROR"))
	if os.Getenv("SERVER_ERROR") == "1" {
		return errors.New("503 error")
	}
	return nil
}