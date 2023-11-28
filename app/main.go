package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jerricodelacruz/goappinsights/appinsights_wrapper"
)

func main() {
	appinsights_wrapper.Init(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	RunFirst()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func RunFirst() {
	client := appinsights_wrapper.NewAppInsightsClient()

	client.StartOperation("OPERATION CORRELATION...")

	client.TrackEvent("FIRST")

	RunSecond()

	RunThird()

	client.EndOperation()
}

func RunSecond() {
	client := appinsights_wrapper.NewAppInsightsClient()

	client.TrackEvent("SECOND")

	RunSecondFirst()
}

func RunSecondFirst() {
	client := appinsights_wrapper.NewAppInsightsClient()

	client.TrackEvent("SECOND FIRST")
}

func RunThird() {
	client := appinsights_wrapper.NewAppInsightsClient()

	client.TrackEvent("THIRD")

	RunSecondFirst()
}
