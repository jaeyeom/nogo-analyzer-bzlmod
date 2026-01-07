// Package analyzer provides individual staticcheck analyzers for nogo.
// The analyzer name is injected at build time via x_defs.
package analyzer

import (
	"golang.org/x/tools/go/analysis"

	"github.com/jaeyeom/nogo-analyzer-bzlmod/staticcheck/util"
)

// name is replaced at build time via x_defs in the BUILD file.
// Each analyzer target injects its own name (e.g., "SA1000", "ST1001").
var name = "dummy value please replace using x_defs"

// Analyzer is the exported analyzer instance for nogo consumption.
// nolint: gochecknoglobals
var Analyzer *analysis.Analyzer = util.FindAnalyzerByName(name)
