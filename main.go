package main

import (
	"log"
	"net/http"

	"github.com/nongdenchet/covidform/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/generate", handler.NewHandler())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
