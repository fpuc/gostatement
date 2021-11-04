package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/fpuc/gostatement"
)

func main() {
	unitchecker.Main(gostatement.Analyzer)
}
