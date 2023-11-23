package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func main() {
	TrackGroupWithHerierchy()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func TrackGroupWithHerierchy() {
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	client.Context().Tags.Operation().SetId("e49635cc-007d-4006-b661-dd23a5946cf5")

	firstEvent := appinsights.NewEventTelemetry("1ST EVENT")
	firstEvent.Properties["property"] = "1STPROPERTY"
	client.Track(firstEvent)

	secondEvent := appinsights.NewEventTelemetry("2ND EVENT SUB OF 1ST EVENT")
	secondEvent.Properties["property"] = "2NDPROPERTY"
	client.Track(secondEvent)

	thirdEvent := appinsights.NewEventTelemetry("3RD EVENT")
	thirdEvent.Properties["property"] = "3RDPROPERTY"
	client.Track(thirdEvent)
}
