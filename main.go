package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"mygra.tech/project1/Config"
	"mygra.tech/project1/Routes"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func main() {
	err := godotenv.Load(".env")
	ctx := context.Background()

	
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	do3Sum()
	
	// setupJaeger()
	
	// Config.InitCassandra()
	
	db := Config.DatabaseOpen()
	
	// Setup routes
	r := Routes.SetupRouter(db)
	
	// Setup port
	serverPort := os.Getenv("SERVER_PORT")
	
	go Config.Produce(ctx)
	Config.Consume(ctx)
	
	// Running
	r.Run(":" + serverPort)
}

func setupJaeger() {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "your_service_name",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
}

func do3Sum() {
	
}
