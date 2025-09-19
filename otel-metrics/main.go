package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// 1. Using context in Go is essential for managing timeouts, cancellation, and request-scoped data across
	// API boundaries and goroutines.
	ctx := context.Background()

	// 1. Create the OTLP HTTP exporter (no TLS)
	exporter, err := otlpmetrichttp.New(
		ctx,
		otlpmetrichttp.WithEndpoint("localhost:9091"), // replace with your OpenTelemetry Collector endpoint
		otlpmetrichttp.WithURLPath("/api/v1/otlp"),    // Set the OTLP HTTP path
		otlpmetrichttp.WithInsecure(),                 // use TLS in production
	)

	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}

	// 2. Configure a PeriodicReader to push metrics every 10s
	reader := sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second))
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	defer mp.Shutdown(ctx)

	otel.SetMeterProvider(mp)

	// Create a counter
	meter := otel.Meter("simple")
	counter, _ := meter.Int64Counter("ping_count") // You will be able to query "ping_count_total" on Prometheus and Grafana UI

	for i := 1; ; i++ {
		counter.Add(ctx, 1)
		log.Printf("ping #%d sent", i)
		time.Sleep(10 * time.Second)
	}
}
