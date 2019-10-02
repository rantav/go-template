package main

import (
	"github.com/rantav/go-template/cmd"
	"github.com/rantav/go-template/log"
	"github.com/rantav/go-template/version"
)

func main() {
	log.Setup(false)
	version.LogVersion()
	cmd.Execute()
}
