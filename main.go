package main

import (
	"go-zh2en/check"
	"go-zh2en/replace"
	"os"
)

func main() {
	params := os.Args
	if len(params) > 1 {
		serverType := params[1]
		dir := params[2]
		if serverType == "check" {
			check.Run(dir)
		}
		if serverType == "replace" {
			replace.TraverseDirectory(dir)
		}
	}
}
