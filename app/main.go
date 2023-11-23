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
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	// Set role instance name globally -- this is usually the
	// name of the service submitting the telemetry
	client.Context().Tags.Cloud().SetRole("my_go_server")

	// Set the role instance to the host name.  Note that this is
	// done automatically by the SDK.
	hostname, _ := os.Hostname()
	client.Context().Tags.Cloud().SetRoleInstance(hostname)

	client.Context().Tags.Operation().SetId("e49635cc-007d-4006-b661-dd23a5946cf8")

	firstEvent := appinsights.NewEventTelemetry("1ST EVENT")
	firstEvent.Properties["property"] = "1STPROPERTY"
	client.Track(firstEvent)

	client.Context().Tags.Operation().SetParentId(firstEvent.Tags.Operation().GetId())
	secondEvent := appinsights.NewEventTelemetry("2ND EVENT SUB OF 1ST EVENT")
	secondEvent.Properties["property"] = "2NDPROPERTY"
	client.Track(secondEvent)

	for k := range client.Context().Tags.Operation() {
		delete(client.Context().Tags.Operation(), k)
	}

	client.Context().Tags.Operation().SetId("e49635cc-007d-4006-b661-dd23a5946cf8")

	thirdEvent := appinsights.NewEventTelemetry("3RD EVENT")
	thirdEvent.Properties["property"] = "2NDPROPERTY"
	client.Track(thirdEvent)
}
