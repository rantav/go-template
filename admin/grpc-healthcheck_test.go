package admin

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"

	grpclib "github.com/rantav/go-template/internal/generated/grpc"
)

type mockClient struct {
	succeed bool // should the check succeed or not?
}

func (c *mockClient) Healthcheck(
	ctx context.Context,
	in *empty.Empty,
	opts ...grpc.CallOption) (
	*grpclib.HealthCheckResponse, error,
) {
	if c.succeed {
		return &grpclib.HealthCheckResponse{
			Status: grpclib.HealthCheckResponse_Serving,
		}, nil
	}
	return nil, fmt.Errorf("FAIL")
}

func TestGrpcHealthcheck_happy(t *testing.T) {
	assert := assert.New(t)
	client := &mockClient{true}
	check := NewGrpcHealthcheck(client, "my grpc check")
	assert.NotNil(check)
	assert.Equal("my grpc check", check.Name())
	res, err := check.Execute()
	assert.NoError(err)
	assert.Equal("Serving", res.(*grpclib.HealthCheckResponse).Status.String())
}

func TestGrpcHealthcheck_sad(t *testing.T) {
	assert := assert.New(t)
	client := &mockClient{false}
	check := NewGrpcHealthcheck(client, "my sad grpc check")
	_, err := check.Execute()
	assert.Error(err)
}
