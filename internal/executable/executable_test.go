package executable_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rios0rios0/dev-toolkit/internal/executable"
)

func TestResolverResolve(t *testing.T) {
	t.Parallel()

	t.Run("should return the absolute path when the tool is on PATH", func(t *testing.T) {
		t.Parallel()
		// given
		resolver := &executable.Resolver{
			LookPath: func(name string) (string, error) { return "/usr/bin/" + name, nil },
		}

		// when
		path, err := resolver.Resolve("git")

		// then
		require.NoError(t, err)
		assert.Equal(t, "/usr/bin/git", path)
	})

	t.Run("should look a tool up only once when it is resolved repeatedly", func(t *testing.T) {
		t.Parallel()
		// given
		lookups := 0
		resolver := &executable.Resolver{
			LookPath: func(name string) (string, error) {
				lookups++
				return "/usr/bin/" + name, nil
			},
		}

		// when
		first, firstErr := resolver.Resolve("git")
		second, secondErr := resolver.Resolve("git")

		// then
		require.NoError(t, firstErr)
		require.NoError(t, secondErr)
		assert.Equal(t, first, second)
		assert.Equal(t, 1, lookups)
	})

	t.Run("should resolve each tool on its own when several are requested", func(t *testing.T) {
		t.Parallel()
		// given
		resolver := &executable.Resolver{
			LookPath: func(name string) (string, error) { return "/opt/homebrew/bin/" + name, nil },
		}

		// when
		git, gitErr := resolver.Resolve("git")
		docker, dockerErr := resolver.Resolve("docker")

		// then
		require.NoError(t, gitErr)
		require.NoError(t, dockerErr)
		assert.Equal(t, "/opt/homebrew/bin/git", git)
		assert.Equal(t, "/opt/homebrew/bin/docker", docker)
	})

	t.Run("should return an error when the tool is not on PATH", func(t *testing.T) {
		t.Parallel()
		// given
		resolver := &executable.Resolver{
			LookPath: func(string) (string, error) { return "", errors.New("executable file not found in $PATH") },
		}

		// when
		_, err := resolver.Resolve("docker")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "docker not found in PATH")
	})

	t.Run("should look a tool up again when an earlier lookup failed", func(t *testing.T) {
		t.Parallel()
		// given
		installed := false
		resolver := &executable.Resolver{
			LookPath: func(name string) (string, error) {
				if !installed {
					return "", errors.New("executable file not found in $PATH")
				}
				return "/usr/local/bin/" + name, nil
			},
		}
		_, firstErr := resolver.Resolve("docker")
		installed = true

		// when
		path, err := resolver.Resolve("docker")

		// then
		require.Error(t, firstErr)
		require.NoError(t, err)
		assert.Equal(t, "/usr/local/bin/docker", path)
	})

	t.Run("should refuse the tool when PATH resolves it to a relative path", func(t *testing.T) {
		t.Parallel()
		// given
		resolver := &executable.Resolver{
			LookPath: func(string) (string, error) { return "./git", nil },
		}

		// when
		_, err := resolver.Resolve("git")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), `git resolved to the relative path "./git"`)
	})
}

func TestResolve(t *testing.T) {
	t.Run("should find the running test binary when its directory is the whole PATH", func(t *testing.T) {
		// given
		self, err := os.Executable()
		require.NoError(t, err)
		t.Setenv("PATH", filepath.Dir(self))

		// when
		path, err := executable.Resolve(filepath.Base(self))

		// then
		require.NoError(t, err)
		assert.True(t, filepath.IsAbs(path), "expected an absolute path, got %q", path)
		assert.Equal(t, filepath.Base(self), filepath.Base(path))
	})

	t.Run("should return an error when nothing on PATH matches", func(t *testing.T) {
		// given
		t.Setenv("PATH", t.TempDir())

		// when
		_, err := executable.Resolve("dev-toolkit-no-such-tool")

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "dev-toolkit-no-such-tool not found in PATH")
	})
}
