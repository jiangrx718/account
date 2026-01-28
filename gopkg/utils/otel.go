package utils

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitOpenTelemetry(ctx context.Context) error {
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return err
	}

	// resources, err := resource.New(
	// 	context.Background(),
	// 	resource.WithAttributes(
	// 		attribute.String("service.name", serviceName),
	// 		attribute.String("library.language", "go"),
	// 	),
	// )
	// if err != nil {
	// 	return err
	// }

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		// sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return nil
}
