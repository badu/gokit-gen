package poc

import (
	"context"
	"github.com/go-kit/kit/log"

	"google.golang.org/grpc/metadata"
)

type metaContext string

const (
	correlationID     metaContext = "correlation-id"
	responseHDR       metaContext = "my-response-header"
	responseTRLR      metaContext = "my-response-trailer"
	correlationIDTRLR metaContext = "correlation-id-consumed"
)

/* client before functions */

func InjectCorrelationID(logger log.Logger) func(context.Context, *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if hdr, ok := ctx.Value(correlationID).(string); ok {
			logger.Log("client_request", "inject", "correlationID", hdr)
			(*md)[string(correlationID)] = append((*md)[string(correlationID)], hdr)
		}
		return ctx
	}
}

func DisplayClientRequestHeaders(logger log.Logger) func(context.Context, *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if len(*md) > 0 {
			logger.Log("client_request", "headers")
			for key, val := range *md {
				logger.Log("client_request", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

/* server before functions */

func extractCorrelationID(logger log.Logger) func(context.Context, metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		if hdr, ok := md[string(correlationID)]; ok {
			cID := hdr[len(hdr)-1]
			ctx = context.WithValue(ctx, correlationID, cID)
			logger.Log("server_request", "extract", "correlationID", cID)
		}
		return ctx
	}
}

func displayServerRequestHeaders(logger log.Logger) func(context.Context, metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		if len(md) > 0 {
			logger.Log("server_request", "headers")
			for key, val := range md {
				logger.Log("server_request", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

/* server after functions */

func injectResponseHeader(logger log.Logger) func(context.Context, *metadata.MD, *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD, _ *metadata.MD) context.Context {
		*md = metadata.Join(*md, metadata.Pairs(string(responseHDR), "has-a-value"))
		return ctx
	}
}

func displayServerResponseHeaders(logger log.Logger) func(context.Context, *metadata.MD, *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD, _ *metadata.MD) context.Context {
		if len(*md) > 0 {
			logger.Log("server_response", "headers")
			for key, val := range *md {
				logger.Log("server_response", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

func injectResponseTrailer(logger log.Logger) func(context.Context, *metadata.MD, *metadata.MD) context.Context {
	return func(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
		*md = metadata.Join(*md, metadata.Pairs(string(responseTRLR), "has-a-value-too"))
		return ctx
	}
}

func injectConsumedCorrelationID(logger log.Logger) func(context.Context, *metadata.MD, *metadata.MD) context.Context {
	return func(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
		if hdr, ok := ctx.Value(correlationID).(string); ok {
			logger.Log("server_response", "inject", "correlationID", hdr)
			*md = metadata.Join(*md, metadata.Pairs(string(correlationIDTRLR), hdr))
		}
		return ctx
	}
}
func displayServerResponseTrailers(logger log.Logger) func(context.Context, *metadata.MD, *metadata.MD) context.Context {
	return func(ctx context.Context, _ *metadata.MD, md *metadata.MD) context.Context {
		if len(*md) > 0 {
			logger.Log("server_response", "trailers")
			for key, val := range *md {
				logger.Log("server_response", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

/* client after functions */

func DisplayClientResponseHeaders(logger log.Logger) func(context.Context, metadata.MD, metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD, _ metadata.MD) context.Context {
		if len(md) > 0 {
			logger.Log("client_response", "headers")
			for key, val := range md {
				logger.Log("client_response", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

func DisplayClientResponseTrailers(logger log.Logger) func(context.Context, metadata.MD, metadata.MD) context.Context {
	return func(ctx context.Context, _ metadata.MD, md metadata.MD) context.Context {
		if len(md) > 0 {
			logger.Log("client_response", "trailers")
			for key, val := range md {
				logger.Log("client_response", "", "key", key, "value", val[len(val)-1])
			}
		}
		return ctx
	}
}

func ExtractConsumedCorrelationID(logger log.Logger) func(context.Context, metadata.MD, metadata.MD) context.Context {
	return func(ctx context.Context, _ metadata.MD, md metadata.MD) context.Context {
		if hdr, ok := md[string(correlationIDTRLR)]; ok {
			logger.Log("client_response", "extract", "correlationID", hdr[len(hdr)-1])
			ctx = context.WithValue(ctx, correlationIDTRLR, hdr[len(hdr)-1])
		}
		return ctx
	}
}

/* CorrelationID context handlers */

func SetCorrelationID(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, correlationID, v)
}
