// Package sparkles exposes the upstream CLI entry point for a multicall binary.
package sparkles

import (
	"context"
	"github.com/cli/cli/v2/internal/build"
	"github.com/cli/cli/v2/internal/ghcmd"
)

// Main runs once per process, using os.Args, the environment and standard streams.
func Main(ctx context.Context) int {
	build.Version = "2.100.0"
	return int(ghcmd.MainContext(ctx))
}
