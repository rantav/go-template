package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup logging
func Setup(useJSON bool) {
	if useJSON {
		log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
