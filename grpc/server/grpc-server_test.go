package grpcserver

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	server, err := New()
	assert.NoError(err)
	assert.NotNil(server)
}

func TestHealthcheck(t *testing.T) {
	assert := assert.New(t)
	server := &server{}
	res, err := server.Healthcheck(context.Background(), &empty.Empty{})
	assert.NoError(err)
	assert.Equal("Serving", res.Status.String())
}
