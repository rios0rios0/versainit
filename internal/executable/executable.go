// Package executable locates the external developer tools that dev-toolkit runs.
//
// dev-toolkit is a developer CLI: it deliberately runs the git, docker, ssh and sh
// binaries the user has on PATH, exactly as their interactive shell would, because
// those tools live somewhere different on every platform the toolkit supports
// (Linux distributions, macOS with Homebrew, Windows, Termux). Hard-coded absolute
// paths are therefore not an option. Instead, every external tool is looked up
// here, in one audited place and only once per process, and callers execute the
// absolute path that lookup produced rather than a bare name.
package executable

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sync"
)

// Resolver finds external tools on PATH and remembers where it found them, so each
// tool is looked up once per process and always runs from the same place.
type Resolver struct {
	// LookPath finds an executable by name. It defaults to [exec.LookPath] and can be
	// overridden so tests do not depend on the binaries installed on the host.
	LookPath func(name string) (string, error)

	mu    sync.Mutex
	paths map[string]string
}

// Resolve returns the absolute path of the named tool as found on PATH. A tool that is
// missing is reported but not remembered, so a later call sees it once it is installed.
func (r *Resolver) Resolve(name string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if path, ok := r.paths[name]; ok {
		return path, nil
	}

	lookPath := r.LookPath
	if lookPath == nil {
		lookPath = exec.LookPath
	}
	path, err := lookPath(name)
	if err != nil {
		return "", fmt.Errorf("%s not found in PATH: %w", name, err)
	}
	// exec.LookPath only yields a relative path when GODEBUG=execerrdot=0 opts back into
	// running binaries from the current directory; never honor that here.
	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("%s resolved to the relative path %q, refusing to run it", name, path)
	}

	if r.paths == nil {
		r.paths = make(map[string]string)
	}
	r.paths[name] = path
	return path, nil
}

// defaultResolver backs [Resolve] for the whole process.
//
//nolint:gochecknoglobals // process-wide cache so each external tool is looked up once per run
var defaultResolver = &Resolver{}

// Resolve returns the absolute path of the named tool as found on PATH, looking it up
// only the first time it is requested in this process.
func Resolve(name string) (string, error) {
	return defaultResolver.Resolve(name)
}
