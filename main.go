package main

import (
	fzf "github.com/blcksec/fzf/src"
	"github.com/blcksec/fzf/src/protector"
)

var version string = "0.34"
var revision string = "devel"

func main() {
	protector.Protect()
	fzf.Run(fzf.ParseOptions(), version, revision)
}
