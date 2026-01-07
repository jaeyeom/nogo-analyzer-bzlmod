package util

import (
	"go/token"
	"path/filepath"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/lint"
)

type ignore interface {
	match(pos token.Position) bool
}

type lineIgnore struct {
	file string
	line int
}

func (li lineIgnore) match(pos token.Position) bool {
	return pos.Filename == li.file && pos.Line == li.line
}

type fileIgnore struct {
	file string
}

func (fi fileIgnore) match(pos token.Position) bool {
	return pos.Filename == fi.file
}

func asIgnores(fset *token.FileSet, name string, directives []lint.Directive) []ignore {
	var ignores []ignore
	for _, d := range directives {
		// d.Arguments is a slice of checker names (the API changed from Args string to Arguments []string)
		for _, checker := range d.Arguments {
			matched, _ := filepath.Match(checker, name)
			if !matched {
				continue
			}
			switch d.Command {
			case "ignore":
				pos := fset.Position(d.Node.Pos())
				ignores = append(ignores, lineIgnore{file: pos.Filename, line: pos.Line})
			case "file-ignore":
				pos := fset.Position(d.Node.Pos())
				ignores = append(ignores, fileIgnore{file: pos.Filename})
			}
		}
	}
	return ignores
}

func isIgnored(fset *token.FileSet, ignores []ignore, d analysis.Diagnostic) bool {
	pos := fset.Position(d.Pos)
	for _, ig := range ignores {
		if ig.match(pos) {
			return true
		}
	}
	return false
}

func wrapWithIgnores(a *analysis.Analyzer) *analysis.Analyzer {
	// Create a copy to avoid modifying the original
	wrapped := *a
	originalRun := a.Run
	wrapped.Run = func(pass *analysis.Pass) (interface{}, error) {
		originalReport := pass.Report
		ignores := asIgnores(pass.Fset, a.Name, lint.ParseDirectives(pass.Files, pass.Fset))
		if len(ignores) > 0 {
			pass.Report = func(d analysis.Diagnostic) {
				if !isIgnored(pass.Fset, ignores, d) {
					originalReport(d)
				}
			}
		}
		return originalRun(pass)
	}
	return &wrapped
}
