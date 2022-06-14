package xtracer

import (
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewJaegerTracer)

type XTracer struct {
	Tracer opentracing.Tracer
	closer io.Closer
	err    error
}

func NewJaegerTracer(vp *viper.Viper) *XTracer {
	cfg := &config.Configuration{
		ServiceName: vp.GetString("TRACER.SERVICENAME"),
		// ALL
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: vp.GetString("TRACER.AGENTADDR"),
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	opentracing.SetGlobalTracer(tracer)
	xtracer := &XTracer{
		Tracer: opentracing.GlobalTracer(),
		closer: closer,
		err:    err,
	}
	return xtracer
}
