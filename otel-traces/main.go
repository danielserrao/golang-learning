package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func main() {
	ctx := context.Background()

	// Create OTLP/HTTP exporter (localhost:4318)
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(), // No TLS for local dev
	)

	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	// Create a tracer provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("my-go-service2"),
		)),
	)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)

	// Create a tracer and send a span
	tracer := otel.Tracer("example.com/trace")

	ctx, span := tracer.Start(ctx, "my-span")
	span.SetAttributes(attribute.String("env", "dev"))
	time.Sleep(500 * time.Millisecond)
	span.End()

	log.Println("Trace sent to Grafana Agent via OTLP HTTP!")
}
