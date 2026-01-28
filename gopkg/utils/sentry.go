package utils

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func InitSentryFromViper() error {

	viper.SetDefault("sentry.dsn", "")
	viper.SetDefault("sentry.enable_tracing", true)
	viper.SetDefault("sentry.traces_sample_rate", 1.0)
	viper.SetDefault("sentry.profiles_sample_rate", 1.0)

	if err := sentry.Init(sentry.ClientOptions{
		Debug:            Debug(),
		Release:          ReleaseTag(),
		Environment:      Env(),
		AttachStacktrace: true,
		Dsn:              viper.GetString("sentry.dsn"),
		// EnableTracing:    viper.GetBool("sentry.enable_tracing"),
		// We recommend adjusting these values in production:
		// TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
		// 	// As an example, this does not send some
		// 	// transactions to Sentry based on their name.
		// 	if ctx.Span.Name == "GET /health" || ctx.Span.Name == "GET /ping" {
		// 		return 0.0
		// 	}
		// 	return viper.GetFloat64("sentry.traces_sample_rate")
		// }),
		// // The sampling rate for profiling is relative to TracesSampleRate:
		// ProfilesSampleRate: viper.GetFloat64("sentry.profiles_sample_rate"),
	}); err != nil {
		return nil
	}

	// tp := trace.NewTracerProvider(
	// 	trace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	// )
	// otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
	return nil
}
