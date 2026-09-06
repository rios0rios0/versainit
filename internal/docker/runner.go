package docker

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/rios0rios0/dev-toolkit/internal/executable"
)

// dockerExecutable is the docker CLI as it has to be found on the user's PATH.
const dockerExecutable = "docker"

// Runner abstracts docker command execution for testability.
type Runner interface {
	Run(args ...string) error
	Output(args ...string) (string, error)
}

// DefaultRunner executes real docker commands via [exec.CommandContext], through the
// docker binary that [executable.Resolve] located on PATH.
type DefaultRunner struct{}

func (r *DefaultRunner) Run(args ...string) error {
	cmd, err := r.command(args...)
	if err != nil {
		return err
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", strings.Join(args, " "), strings.TrimSpace(string(output)))
	}
	return nil
}

func (r *DefaultRunner) Output(args ...string) (string, error) {
	cmd, err := r.command(args...)
	if err != nil {
		return "", err
	}
	output, err := cmd.Output()
	if err != nil {
		exitErr := &exec.ExitError{}
		if errors.As(err, &exitErr) {
			stderr := strings.TrimSpace(string(exitErr.Stderr))
			return "", fmt.Errorf("%s: %s", strings.Join(args, " "), stderr)
		}
		msg := strings.TrimSpace(string(output))
		if msg == "" {
			return "", fmt.Errorf("%s: %w", strings.Join(args, " "), err)
		}
		return "", fmt.Errorf("%s: %s", strings.Join(args, " "), msg)
	}
	return strings.TrimSpace(string(output)), nil
}

// command prepares a docker invocation that runs the resolved binary by absolute path.
func (r *DefaultRunner) command(args ...string) (*exec.Cmd, error) {
	docker, err := executable.Resolve(dockerExecutable)
	if err != nil {
		return nil, err
	}
	cmd := exec.CommandContext(context.Background(), docker, args...) // #nosec G204
	cmd.Stdin = nil
	return cmd, nil
}
