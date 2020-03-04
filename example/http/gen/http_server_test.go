package factory_test

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/badu/gokit-gen/example/http/gen"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	kitHTTP "github.com/go-kit/kit/transport/http"
)

type HTTPMockedRepository struct{}

func (r *HTTPMockedRepository) MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error) {
	// TODO : implement me
	return 20, 220, nil
}

func (r *HTTPMockedRepository) GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error) {
	// TODO : implement me
	return nil, nil
}

func (r *HTTPMockedRepository) Status(ctx context.Context, page int32, perPage int32) (string, bool, error) {
	// TODO : implement me
	return "blah blah", true, nil
}

// classic grpc call for MakeBox
func TestHTTPMakeBox(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	service := NewService(logger, &HTTPMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(service, []endpoint.Middleware{}) // authMiddleware

	options := []kitHTTP.ServerOption{
		kitHTTP.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	server := httptest.NewServer(MakeMakeBoxHTTPHandler(endpoints.MakeBoxEndpoint, logger, options...))
	defer server.Close()

	ctx := context.Background()
	var clientOptions []kitHTTP.ClientOption

	client, err := NewHTTPClient(server.URL, logger, clientOptions...)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	// TODO : load the payloads
	page, perPage, err := client.MakeBox(ctx, "", 0, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("page : %v", page)
	t.Logf("perPage : %v", perPage)
}

// classic grpc call for GetBoxes
func TestHTTPGetBoxes(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	service := NewService(logger, &HTTPMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(service, []endpoint.Middleware{}) // authMiddleware

	options := []kitHTTP.ServerOption{
		kitHTTP.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	server := httptest.NewServer(MakeGetBoxesHTTPHandler(endpoints.GetBoxesEndpoint, logger, options...))
	defer server.Close()

	ctx := context.Background()
	var clientOptions []kitHTTP.ClientOption

	client, err := NewHTTPClient(server.URL, logger, clientOptions...)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	// TODO : load the payloads
	boxes, err := client.GetBoxes(ctx, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("boxes : %v", boxes)
}

// classic grpc call for Status
func TestHTTPStatus(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	service := NewService(logger, &HTTPMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(service, []endpoint.Middleware{}) // authMiddleware

	options := []kitHTTP.ServerOption{
		kitHTTP.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	server := httptest.NewServer(MakeStatusHTTPHandler(endpoints.StatusEndpoint, logger, options...))
	defer server.Close()

	ctx := context.Background()
	var clientOptions []kitHTTP.ClientOption

	client, err := NewHTTPClient(server.URL, logger, clientOptions...)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	// TODO : load the payloads
	serviceName, ok, err := client.Status(ctx, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("serviceName : %v", serviceName)
	t.Logf("ok : %v", ok)
}
