package utils

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

// InitTracer return new opentracer, closer and error with given service name and jaeger agent host
func InitTracer(service, jaegerAgentHost string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: jaegerAgentHost,
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger), jaegercfg.Metrics(metrics.NullFactory))
	if err != nil {
		return nil, nil, err
	}

	return tracer, closer, nil
}
