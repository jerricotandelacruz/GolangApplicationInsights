package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
)

func main() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	// TELEMETRY SUBMISSION

	// TRACK EVENT
	client.TrackEvent("TRACK EVENT:\n Log a user action with the specified name")

	// TRACK METRIC
	client.TrackMetric("TRACK METRIC:\n Log a numeric value that is not specified with a specific event.\n Typically used to send regular reports of performance indicators.", 999.999)

	// TRACK TRACE
	client.TrackTrace("TRACK TRACE (VERBOSE):\n Log a trace message with the specified severity level.", contracts.Verbose)
	client.TrackTrace("TRACK TRACE (INFORMATION):\n Log a trace message with the specified severity level.", contracts.Information)
	client.TrackTrace("TRACK TRACE (WARNING):\n Log a trace message with the specified severity level.", contracts.Warning)
	client.TrackTrace("TRACK TRACE (ERROR):\n Log a trace message with the specified severity level.", contracts.Error)
	client.TrackTrace("TRACK TRACE (CRITICAL):\n Log a trace message with the specified severity level.", contracts.Critical)
	client.TrackTrace("TRACK TRACE (UNKNOWN):\n Log a trace message with the specified severity level.", 5) // UNKNOWN

	// TRACT REQUEST
	var duration time.Duration = 10
	client.TrackRequest("GET", "https://example.com", duration, "200")

	// TRACK REMOTEE DEPENDENCY
	client.TrackRemoteDependency("TRACK REMOTE DEPENDENCY (SUCCESS):\n Log a dependency with the specified name, type, target, and success status.", "DEPENDENCY TYPE", "TARGET", true)
	client.TrackRemoteDependency("TRACK REMOTE DEPENDENCY (FAILED):\n Log a dependency with the specified name, type, target, and success status.", "DEPENDENCY TYPE", "TARGET", false)

	// TRACK AVAILABILITY
	client.TrackAvailability("TRACK AVAILABILITY (SUCCESS):\n Log an availability test result with the specified test name,\n duration, and success status.\n AVAILABILITY NAME", duration, true)
	client.TrackAvailability("TRACK AVAILABILITY (FAILED):\n Log an availability test result with the specified test name,\n duration, and success status.\n AVAILABITLITY NAME", duration, false)

	// TRACK EXCEPTIONS
	dummyException := struct {
		Description  string
		ErrorMessage string
	}{
		Description:  "TRACK EXCEPTION (OBJECT):\n Log an exception with the specified error, which may be a string,\n error or Stringer. The current callstack is collected\n automatically.",
		ErrorMessage: "DUMMY ERROR MESSAGE BLAH BLAH",
	}

	client.TrackException(dummyException)
	client.TrackException("TRACKT EXCEPTION (STRING)")

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
