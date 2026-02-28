# Contributing

Contributions are welcome. By participating, you agree to maintain a respectful and constructive environment.

For coding standards, testing patterns, architecture guidelines, commit conventions, and all
development practices, refer to the **[Development Guide](https://github.com/rios0rios0/guide/wiki)**.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26+
- [Make](https://www.gnu.org/software/make/)
- [Git](https://git-scm.com/) 2.0+

## Development Workflow

1. Fork and clone the repository
2. Create a branch: `git checkout -b feat/my-change`
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Build the binary:
   ```bash
   make build
   ```
   This compiles the CLI to `bin/vinit` from `./cmd/versainit`.
5. Run the application (without building):
   ```bash
   make run
   ```
6. Build a debug binary (with symbols for debuggers):
   ```bash
   make debug
   ```
7. Install locally:
   ```bash
   make install
   ```
   This builds and copies `bin/vinit` to `/usr/local/bin/vinit`.
8. Run tests:
   ```bash
   go test ./...
   ```
9. Commit following the [commit conventions](https://github.com/rios0rios0/guide/wiki/Life-Cycle/Git-Flow)
10. Open a pull request against `main`
