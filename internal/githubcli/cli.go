package githubcli

import (
	"bytes"
	"github.com/cli/go-gh"
)

type GhCliExecutor interface {
	Exec(args ...string) (stdOut, stdErr bytes.Buffer, err error)
}

type GhCli struct {
}

func (c GhCli) Exec(args ...string) (stdOut, stdErr bytes.Buffer, err error) {
	return gh.Exec(args...)
}

func NewGhCli() GhCli {
	return GhCli{}
}
