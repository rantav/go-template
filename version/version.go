package version

import (
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	ServiceName = "go-template"
)

// All these values will be replaced by actual values at build time.
var (
	GitHash          = "-"
	GitBranch        = "-"
	GitTag           = "-"
	GitCommitMessage = "-"
	BuildTime        = "-"
)

func init() {
	// Since there's an issue passing strings with spaces to ldflags,
	// we replace the strings to _ at build time and here we replace them back.
	GitCommitMessage = strings.Replace(GitCommitMessage, "_", " ", -1)
}

// LogVersion simply prints the current version information to the log
func LogVersion() {
	log.Info().
		Str("git-hash", GitHash).
		Str("git-message", GitCommitMessage).
		Str("git-branch", GitBranch).
		Str("git-tag", GitTag).
		Str("build-time", BuildTime).
		Msgf(`Version information:
	git-hash: %s
	git-message: %s
	git-branch: %s
	git-tag: %s
	build-time: %s
	`,
			GitHash, GitCommitMessage, GitBranch, GitTag, BuildTime)
}
