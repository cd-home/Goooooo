package tracer

import (
	"context"

	"github.com/GodYao1995/Goooooo/pkg/xtracer"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing(xtracer *xtracer.XTracer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var newCtxWithSpan context.Context
		var span opentracing.Span
		spanCtx, err := xtracer.Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err != nil {
			span, newCtxWithSpan = opentracing.StartSpanFromContextWithTracer(ctx.Request.Context(), xtracer.Tracer, ctx.Request.URL.Path)
		} else {
			span, newCtxWithSpan = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				xtracer.Tracer,
				ctx.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}
		defer span.Finish()

		var traceID string
		var spanID string
		var spanContext = span.Context()
		jaegerContext, ok := spanContext.(jaeger.SpanContext)
		if ok {
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
			ctx.Set("X-Trace-ID", traceID)
			ctx.Set("X-Span-ID", spanID)
			ctx.Writer.Header().Set("X-Trace-ID", traceID)
		}
		// 放到 request 的 ctx
		ctx.Request = ctx.Request.WithContext(newCtxWithSpan)
		ctx.Next()
	}
}
