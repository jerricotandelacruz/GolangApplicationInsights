package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func main() {
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))
	client.TrackEvent("App Connected")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
