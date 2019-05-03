// +build tools

package fit

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/jgautheron/goconst/cmd/goconst"
	_ "github.com/kisielk/errcheck"
	_ "github.com/mdempsky/unconvert"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "mvdan.cc/gofumpt/gofumports"
)
