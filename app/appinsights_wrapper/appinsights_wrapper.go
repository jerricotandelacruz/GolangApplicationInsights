package appinsights_wrapper

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

// AppInsightsClient holds the configuration and the AI Telemetry client.
type AppInsightsClient struct {
	InstrumentationKey string
	Client             appinsights.TelemetryClient
}

// NewAppInsightsClient creates a new AppInsightsClient with the given instrumentation key.
func NewAppInsightsClient(instrumentationKey string) *AppInsightsClient {
	client := appinsights.NewTelemetryClient(instrumentationKey)
	return &AppInsightsClient{
		InstrumentationKey: instrumentationKey,
		Client:             client,
	}
}

// SetInstrumentationKey sets the instrumentation key for the AppInsightsClient.
func (c *AppInsightsClient) SetInstrumentationKey(instrumentationKey string) {
	c.InstrumentationKey = instrumentationKey
	c.Client = appinsights.NewTelemetryClient(instrumentationKey)
}

// Other functions, methods, and types can be added here based on your requirements.
