package dummy

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
)

func TrackBasicTransactions() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	// TELEMETRY SUBMISSION

	// TRACK EVENT
	client.TrackEvent("[BASIC]TRACK EVENT:\n Log a user action with the specified name")

	// TRACK METRIC
	client.TrackMetric("[BASIC]TRACK METRIC:\n Log a numeric value that is not specified with a specific event.\n Typically used to send regular reports of performance indicators.", 999.999)

	// TRACK TRACE
	client.TrackTrace("[BASIC]TRACK TRACE (VERBOSE):\n Log a trace message with the specified severity level.", contracts.Verbose)
	client.TrackTrace("[BASIC]TRACK TRACE (INFORMATION):\n Log a trace message with the specified severity level.", contracts.Information)
	client.TrackTrace("[BASIC]TRACK TRACE (WARNING):\n Log a trace message with the specified severity level.", contracts.Warning)
	client.TrackTrace("[BASIC]TRACK TRACE (ERROR):\n Log a trace message with the specified severity level.", contracts.Error)
	client.TrackTrace("[BASIC]TRACK TRACE (CRITICAL):\n Log a trace message with the specified severity level.", contracts.Critical)
	client.TrackTrace("[BASIC]TRACK TRACE (UNKNOWN):\n Log a trace message with the specified severity level.", 5) // UNKNOWN

	// TRACT REQUEST
	var duration time.Duration = 10
	client.TrackRequest("GET", "https://basic-example.com", duration, "200")

	// TRACK REMOTE DEPENDENCY
	client.TrackRemoteDependency("[BASIC]TRACK REMOTE DEPENDENCY (SUCCESS):\n Log a dependency with the specified name, type, target, and success status.", "DEPENDENCY TYPE", "TARGET", true)
	client.TrackRemoteDependency("[BASIC]TRACK REMOTE DEPENDENCY (FAILED):\n Log a dependency with the specified name, type, target, and success status.", "DEPENDENCY TYPE", "TARGET", false)

	// TRACK AVAILABILITY
	client.TrackAvailability("[BASIC]TRACK AVAILABILITY (SUCCESS):\n Log an availability test result with the specified test name,\n duration, and success status.\n AVAILABILITY NAME", duration, true)
	client.TrackAvailability("[BASIC]TRACK AVAILABILITY (FAILED):\n Log an availability test result with the specified test name,\n duration, and success status.\n AVAILABITLITY NAME", duration, false)

	// TRACK EXCEPTIONS
	dummyException := struct {
		Description  string
		ErrorMessage string
	}{
		Description:  "[BASIC]TRACK EXCEPTION (OBJECT):\n Log an exception with the specified error, which may be a string,\n error or Stringer. The current callstack is collected\n automatically.",
		ErrorMessage: "DUMMY ERROR MESSAGE BLAH BLAH",
	}

	client.TrackException(dummyException)
	client.TrackException("[BASIC]TRACK EXCEPTION (STRING)")
}

func TrackEvent() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	event := appinsights.NewEventTelemetry("TRACK EVENT")
	event.Properties["property"] = "SAMPLE EVENT PROPERTY"

	client.Track(event)
}

func TrackMetric() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	metric := appinsights.NewMetricTelemetry("TRACK METRIC:", 999.999)
	metric.Properties["Queue name"] = "SAMPLE PROPERTY [PROPERTY NAME: Queue name, VALUE: queue name]"
	client.Track(metric)
}

func TrackTrace() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	trace := appinsights.NewTraceTelemetry("TRACK TRACE (VERBOSE):", appinsights.Verbose)

	// You can set custom properties on traces
	trace.Properties["module"] = "SAMPLE PROPERTY [PROPERTY NAME: module, VALUE: server]"

	// You can also fudge the timestamp:
	trace.Timestamp = time.Now()

	// Finally, track it
	client.Track(trace)
}

func TrackRequest() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	var duration time.Duration = 10
	request := appinsights.NewRequestTelemetry("GET", "https://example.com", duration, "200")

	// Note that the timestamp will be set to time.Now() minus the
	// specified duration.  This can be overridden by either manually
	// setting the Timestamp and Duration fields, or with MarkTime:
	requestStartTime := time.Now().Add(-duration)
	requestEndTime := time.Now()
	request.MarkTime(requestStartTime, requestEndTime)

	// Source of request
	request.Source = "https://example-source.com"

	// Success is normally inferred from the responseCode, but can be overridden:
	request.Success = true

	// Request ID's are randomly generated GUIDs, but this can also be overridden:
	request.Id = "REQUIREDID20231120"

	// Custom properties and measurements can be set here
	request.Properties["user-agent"] = "SAMPLE PROPERTY [PROPERTY NAME: user-agent, VALUE: request headers user-agent]"
	request.Measurements["POST size"] = float64(999.999)

	// Context tags become more useful here as well
	request.Tags.Session().SetId("SESSIONID20231120")
	request.Tags.User().SetAccountId("USERID20231120")

	// Finally track it
	client.Track(request)
}

func TrackDependency() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	dependency := appinsights.NewRemoteDependencyTelemetry("Redis cache", "Redis", "<target>", true /* success */)

	// The result code is typically an error code or response status code
	dependency.ResultCode = "OK"

	// Id's can be used for correlation if the remote end is also logging
	// telemetry through application insights.
	dependency.Id = "<request id>"

	// Data may contain the exact URL hit or SQL statements
	dependency.Data = "MGET <args>"

	// The duration can be set directly:
	var duration time.Duration = 10
	dependency.Duration = duration
	// or via MarkTime:
	requestStartTime := time.Now().Add(-duration)
	requestEndTime := time.Now()
	dependency.MarkTime(requestStartTime, requestEndTime)

	// Properties and measurements may be set.
	dependency.Properties["shard-instance"] = "<name>"
	dependency.Measurements["data received"] = float64(999.999)

	// Submit the telemetry
	client.Track(dependency)
}

func TrackException() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	err := errors.New("TRACT EXCEPTION:\n with warning severity level")
	if err != nil {
		exception := appinsights.NewExceptionTelemetry(err)

		// Set the severity level -- perhaps this isn't a critical
		// issue, but we'd *really rather* it didn't fail:
		exception.SeverityLevel = appinsights.Warning

		// One could tweak the number of stack frames to skip by
		// reassigning the callstack -- for instance, if you were to
		// log this exception in a helper method.
		exception.Frames = appinsights.GetCallstack(3 /* frames to skip */)

		// Properties are available as usual
		exception.Properties["input"] = "SAMPLE PROPERTY [PROPERTY NAME: input, VALUE: argument]"

		// Track the exception
		client.Track(exception)
	}
}

func TrackAvailability() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	var callDuration time.Duration = 10
	availability := appinsights.NewAvailabilityTelemetry("test name", callDuration, true /* success */)

	// The run location indicates where the test was run from
	availability.RunLocation = "Phoenix"

	// Diagnostics message
	availability.Message = "DIAGNOSE MESSAGE"
	// Id is used for correlation with the target service
	availability.Id = "REQUEST ID"

	// Timestamp and duration can be changed through MarkTime, similar
	// to other telemetry types with Duration's

	testStartTime := time.Now().Add(-callDuration)
	testEndTime := time.Now()
	availability.MarkTime(testStartTime, testEndTime)

	// Submit the telemetry
	client.Track(availability)
}

func TrackPageView() {
	// SET INSTRUMENTATION KEY
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	pageview := appinsights.NewPageViewTelemetry("Event name", "http://testuri.org/page")

	// A duration is available here.
	pageview.Duration = time.Minute

	// As are the usual Properties and Measurements...

	// Track
	client.Track(pageview)
}

func TrackGroupEvent() {
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
	client.Context().Tags.Operation().SetParentId("e49635cc-007d-4006-b661-dd23a5946cf8")

	// Make a request to fiddle with the telemetry's context
	req := appinsights.NewRequestTelemetry("GET", "http://server/path", time.Millisecond, "200")

	// Set the account ID context tag, for this telemetry item
	// only.  The following are equivalent:
	req.Tags.User().SetAccountId("<user account retrieved from request>")
	req.Tags[contracts.UserAccountId] = "<user account retrieved from request>"

	// This request will have all context tags above.
	client.Track(req)

	var duration time.Duration = 10
	request := appinsights.NewRequestTelemetry("GET", "https://example.com", duration, "200")

	// Note that the timestamp will be set to time.Now() minus the
	// specified duration.  This can be overridden by either manually
	// setting the Timestamp and Duration fields, or with MarkTime:
	requestStartTime := time.Now().Add(-duration)
	requestEndTime := time.Now()
	request.MarkTime(requestStartTime, requestEndTime)

	// Source of request
	request.Source = "https://example-source.com"

	// Success is normally inferred from the responseCode, but can be overridden:
	request.Success = true

	// Request ID's are randomly generated GUIDs, but this can also be overridden:
	request.Id = "REQUIREDID20231120"

	// Custom properties and measurements can be set here
	request.Properties["user-agent"] = "SAMPLE PROPERTY [PROPERTY NAME: user-agent, VALUE: request headers user-agent]"
	request.Measurements["POST size"] = float64(999.999)

	// Context tags become more useful here as well
	request.Tags.Session().SetId("SESSIONID20231120")
	request.Tags.User().SetAccountId("USERID20231120")

	// Finally track it
	client.Track(request)
}

func RunTimeOutFirstSample() {
	fmt.Println("OUT FIRST")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println("PC : ", pc)
	fmt.Println("FILE : ", file)
	fmt.Println("LINE : ", line)
	fmt.Println("OK : ", ok)
}
