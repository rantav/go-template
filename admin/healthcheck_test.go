package admin

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHealthchecks(t *testing.T) {
	assert := assert.New(t)
	checks := CreateHealthchecks(
		// BEGIN __INCLUDE_GRPC__
		&net.TCPAddr{},
		// END __INCLUDE_GRPC__
	)
	assert.NotNil(checks)
}
