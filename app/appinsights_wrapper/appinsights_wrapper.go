package appinsights_wrapper

import (
	"fmt"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var (
	tc *appinsights.TelemetryConfiguration
	c  *telemetryClient
)

type telemetryClient struct {
	appinsights.TelemetryClient
}

func Init(instrumentationKey string) {
	tc = appinsights.NewTelemetryConfiguration(instrumentationKey)
	c = &telemetryClient{
		TelemetryClient: appinsights.NewTelemetryClientFromConfig(tc),
	}
}

func Client() *telemetryClient {
	return c
}

func (c *telemetryClient) StartOperation(name string) {
	c.Context().Tags.Operation().SetId(newUUID().String())
	c.Context().Tags.Operation().SetName(name)
	fmt.Printf("START OPERATION | ID:%s", c.Context().Tags.Operation().GetId())
}

func (c *telemetryClient) EndOperation() {
	fmt.Printf("END OPERATION | ID:%s", c.Context().Tags.Operation().GetId())
	for k := range c.Context().Tags.Operation() {
		delete(c.Context().Tags.Operation(), k)
	}
}
