// +build tools

package main

import (
	// build/test.mk
	_ "github.com/stretchr/testify/assert"

	// build/lint.mk
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/fzipp/gocyclo"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/remyoudompheng/go-misc/deadcode"
	_ "github.com/timakin/bodyclose"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/godoc"
)
