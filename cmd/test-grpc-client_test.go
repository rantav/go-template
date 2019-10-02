package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcClientCmd(t *testing.T) {
	assert := assert.New(t)
	assert.Contains(rootCmd.Commands(), testGrpcClientCmd)
}
