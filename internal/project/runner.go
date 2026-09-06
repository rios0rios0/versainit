package project

import (
	"context"
	"errors"
	"io"
	"os/exec"

	"github.com/rios0rios0/dev-toolkit/internal/executable"
)

// shellExecutable is the POSIX shell that runs project commands, as it has to be found
// on the user's PATH.
const shellExecutable = "sh"

// CommandRunner executes shell commands with passthrough I/O for interactive usage.
type CommandRunner interface {
	RunInteractive(dir, command string) error
}

// DefaultCommandRunner executes real shell commands with passthrough I/O.
type DefaultCommandRunner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (r *DefaultCommandRunner) RunInteractive(dir, command string) error {
	if command == "" {
		return errors.New("empty command")
	}
	shell, err := executable.Resolve(shellExecutable)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(context.Background(), shell, "-c", command) // #nosec G204
	cmd.Dir = dir
	cmd.Stdin = r.Stdin
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr
	return cmd.Run()
}
