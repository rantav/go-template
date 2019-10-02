package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeCmd(t *testing.T) {
	assert := assert.New(t)
	assert.Contains(rootCmd.Commands(), serveCmd)
}
