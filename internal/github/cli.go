package github

import (
	"bytes"
	"github.com/cli/go-gh"
)

type Cli struct {
}

func (c Cli) Exec(args ...string) (stdOut, stdErr bytes.Buffer, err error) {
	return gh.Exec(args...)
}

func NewCli() *Cli {
	return &Cli{}
}
