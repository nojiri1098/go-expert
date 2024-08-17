package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"nojiri1098/go-expert/chapter1/1.1/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
