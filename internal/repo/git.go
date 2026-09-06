package repo

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rios0rios0/dev-toolkit/internal/executable"
)

// gitExecutable is the git CLI as it has to be found on the user's PATH.
const gitExecutable = "git"

// GitRunner abstracts git command execution for testability.
type GitRunner interface {
	Run(dir string, args ...string) error
	Output(dir string, args ...string) string
	Clone(url, target string) error
}

// DefaultGitRunner executes real git commands via [exec.CommandContext], through the
// git binary that [executable.Resolve] located on PATH.
type DefaultGitRunner struct{}

func (r *DefaultGitRunner) Run(dir string, args ...string) error {
	cmd, err := r.command(dir, args...)
	if err != nil {
		return err
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", strings.Join(args, " "), strings.TrimSpace(string(output)))
	}
	return nil
}

func (r *DefaultGitRunner) Output(dir string, args ...string) string {
	cmd, err := r.command(dir, args...)
	if err != nil {
		return ""
	}
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func (r *DefaultGitRunner) Clone(url, target string) error {
	parentDir := filepath.Dir(target)
	// A directory needs the owner execute bit, so 0o700 (not the rule's 0o600 file threshold)
	// is the least-privilege mode, and owner-only is sufficient.
	// nosemgrep: go.lang.correctness.permissions.file_permission.incorrect-default-permission
	if mkdirErr := os.MkdirAll(parentDir, 0o700); mkdirErr != nil {
		return fmt.Errorf("failed to create clone parent directory %s: %w", parentDir, mkdirErr)
	}

	cmd, err := r.command("", "clone", url, target)
	if err != nil {
		return err
	}
	cmd.Env = append(os.Environ(),
		"GIT_SSH_COMMAND=ssh -o StrictHostKeyChecking=accept-new -o BatchMode=yes",
	)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", strings.TrimSpace(string(cmdOutput)))
	}
	return nil
}

// command prepares a git invocation in dir that runs the resolved binary by absolute
// path. An empty dir keeps the current working directory.
func (r *DefaultGitRunner) command(dir string, args ...string) (*exec.Cmd, error) {
	git, err := executable.Resolve(gitExecutable)
	if err != nil {
		return nil, err
	}
	cmd := exec.CommandContext(context.Background(), git, args...) // #nosec G204
	cmd.Dir = dir
	cmd.Stdin = nil
	return cmd, nil
}
