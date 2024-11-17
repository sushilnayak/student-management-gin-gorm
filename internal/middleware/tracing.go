package middleware

import (
	"fmt"
	"io"
	"student-service/internal/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitTracer(cfg *config.Config) (opentracing.Tracer, io.Closer, error) {

	jaegerCfg := jaegercfg.Configuration{
		ServiceName: cfg.Tracing.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: cfg.Tracing.Endpoint,
		},
	}

	tracer, closer, err := jaegerCfg.NewTracer()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
